'use client'
import Image from "next/image";
import {useApiCall} from "./components/Api.jsx"

export default function Home() {
  const {data, loading, error} = useApiCall('https://fake-json-api.mock.beeceptor.com/users')



  return (
    <div className="font-sans grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20">
        <h1>LOGIN</h1>
        <div className="bg-white p-8 rounded-lg shadow-lg">
        <input  name="" className="w-full px-4 py-2 border-black text-black border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-20" id="" placeholder="Usuario" />
        <br />
        <input  name="" className="w-full px-4 py-2 border-black text-black border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-20" id="" placeholder="Password" type="password" />

        </div>
        
    </div>
  );
}
