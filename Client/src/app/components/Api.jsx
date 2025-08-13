'use client'

import { useState, useEffect } from 'react';



// 1. Hook Personalizado para la consulta a la API
// Este hook tiene la única responsabilidad de manejar la lógica de la petición.
// 'url' es la dirección de la API a la que queremos consultar.
export const useApiCall = (url) => {
  // 'data' guardará la información que recibimos de la API.
  const [data, setData] = useState(null);
  // 'loading' indica si la petición está en curso (true) o ha terminado (false).
  const [loading, setLoading] = useState(true);
  // 'error' guardará cualquier mensaje de error que pueda ocurrir.
  const [error, setError] = useState(null);

  // 'useEffect' se utiliza para ejecutar la lógica de la consulta cuando el componente
  // se carga por primera vez o cuando la URL cambia.
  useEffect(() => {
    // Definimos una función asíncrona para manejar la petición.
    const fetchData = async () => {
      try {
        // Marcamos el estado como "cargando" y limpiamos errores anteriores.
        setLoading(true);
        setError(null);

        // Realizamos la petición a la API.
        const response = await fetch(url);
        
        // Verificamos si la respuesta es exitosa. Si no lo es, lanzamos un error.
        if (!response.ok) {
          throw new Error(`Error HTTP: ${response.status}`);
        }
        
        // Convertimos la respuesta a formato JSON.
        const result = await response.json();
        
        // Actualizamos el estado con los datos recibidos.
        setData(result);
      } catch (e) {
        // Si ocurre un error, lo guardamos en el estado.
        setError(e.message);
        setData(null); // Aseguramos que no haya datos si hay un error.
      } finally {
        // Siempre, al final, marcamos 'loading' como false, sin importar el resultado.
        setLoading(false);
      }
    };

    // Llamamos a la función para iniciar la petición.
    fetchData();

    // El array de dependencias '[url]' asegura que el efecto se vuelva a ejecutar
    // solo si la URL de la API cambia.
  }, [url]);

  // El hook devuelve un objeto con los estados 'data', 'loading' y 'error'.
  // Esto hace que sea fácil de usar en cualquier componente.
  return { data, loading, error };
};

