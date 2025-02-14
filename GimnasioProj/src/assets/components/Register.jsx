import '../styles/register.css'
import axios from 'axios';
import React, { useState } from 'react';
import imagenLogin from '../images/imagenLogin.png'

import {useNavigate } from "react-router-dom"
import { MdAlternateEmail } from "react-icons/md";
import { FaUser } from "react-icons/fa";
import { Toaster, toast } from 'react-hot-toast';


function Register () {

    const navigate = useNavigate()
    const [email, setEmail] = useState('');
    const [nombre, setNombre] = useState('');
    const [error, setError] = useState('');
    const [pagoMensual, setPagoMensual] = useState('');
    const [tipoUsuario, setTipoUsuario] = useState(2);

    const handleRegister = async (e) => {
        e.preventDefault();
        setError('');
           
        console.log("Datos antes de enviar:", {
            nombre,
            email,
            pagoMensual: checked,
            tipoUsuario: 2 
        });

         try {
             const response = await axios.post('http://localhost:8080/registro', {
                nombre: nombre,
                email: email,
                pago_mensual: checked,
                tipo_usuario: 2
                
            });
            console.log("Respuesta del backend:", response.data);

            localStorage.setItem('token', response.data.token);
           
            toast.success('Registro exitoso!', {
                    duration: 3000, 
            });
            limpiarCampos()
            navigate("/Registro")

            } catch (err) {
                setError('Correo o contraseña incorrectos');
                toast.error('Correo o contraseña incorrectos', {
                    duration: 3000, 
                });
            }
    };

    const limpiarCampos = () => {
        setEmail(''),
        setNombre(''),
        setChecked(false)
    }

    const [checked, setChecked] = useState(false);

    const handleToggle = () => {
        setChecked(!checked);
    };

    return (
        <>
           <div className="registerWrapper">
                           <div className="registerContainer">
                               <div className="infoRegister">
                                   <form>
                                       <h1>REGISTRO</h1>
                                       <div className='inputs-form'>
                                           <div className='input-container'>
                                               <input 
                                                   type='text' 
                                                   placeholder='Nombre' 
                                                   value={nombre}
                                                   onChange={(e) => setNombre(e.target.value)}
                                                   required
                                                   />
                                               <FaUser className='icon' />
                                           </div>
                                           <div className='input-container'>
                                               <input 
                                                   type='text' 
                                                   placeholder='Correo' 
                                                   value={email}
                                                   onChange={(e) => setEmail(e.target.value)}
                                                   required
                                                />
                                               <MdAlternateEmail className='icon' />
                                              
                                            </div>
                                            <div className='input-container'>
                                                <span className="switch-label">{checked ? "Pago Mensual Activo" : "Pago Mensual Inactivo"}</span>
                                                <label className='switch'>
                                                    <input 
                                                        type='checkbox' 
                                                        checked={checked} 
                                                        value={pagoMensual}
                                                        onChange={(e) => {
                                                            handleToggle();
                                                            setPagoMensual(e.target.checked); 
                                                        }} 
                                                       
                                                    />
                                                    <span className='slider'></span>
                                                </label>
                                            </div>
                                       </div>
                                      
                                       <button onClick={handleRegister}>Registrarse</button>
           
                                       <div className="new-account">
                                           <p>¿Ya tienes cuenta? <a href="/">Iniciar Sesion</a></p>
                                       </div>
                                   </form>
                               </div>
                               <div className="imgLogin">
                                   <img src={imagenLogin} alt="Imagen Gimnasio" />
                               </div>
                           </div>
                       </div>
    
        </>

    )
}
export default Register