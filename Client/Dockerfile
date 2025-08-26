# ETAPA 1: Construcción
# Utiliza una imagen oficial de Node.js como base para la etapa de construcción.
FROM node:24-alpine3.21 AS builder

# Establece el directorio de trabajo dentro del contenedor.
WORKDIR /app

# Copia los archivos de manifiesto de paquetes para aprovechar el caché de Docker.
# Solo los cambios en estos archivos desencadenarán una nueva instalación de dependencias.
COPY package.json package-lock.json* ./

# Instala todas las dependencias del proyecto.
RUN npm ci

# Copia el resto del código fuente a la imagen de construcción.
COPY . .

# Construye la aplicación de Next.js para producción.
# Esto genera el código estático y optimizado.
RUN npm run build

# ETAPA 2: Ejecución
# Usa una imagen de Node.js más pequeña y minimalista para la ejecución.
# Esto reduce el tamaño final del contenedor y el riesgo de seguridad.
FROM node:24-alpine3.21

# Establece el directorio de trabajo.
WORKDIR /app

# Copia los archivos de configuración de Next.js.
COPY --from=builder /app/package.json ./

# Copia los archivos de producción de la etapa de construcción.
# .next es la carpeta de salida de la construcción.
# public es la carpeta de activos estáticos.
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

# Opcional: Si tienes un archivo `next.config.js`, cópialo también.
COPY --from=builder /app/next.config.mjs ./

#COPY --from=builder /app/node_modules ./node_modules

# Expone el puerto por defecto de Next.js
EXPOSE 3000

# Comando para iniciar la aplicación en modo de producción.
CMD ["npm", "run", "start"]
