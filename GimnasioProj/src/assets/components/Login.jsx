import imagenLogin from '../images/imagenLogin.png'
import '../styles/login.css'
import React, { useState } from 'react';
import axios from 'axios';

import {useNavigate } from "react-router-dom"
import { MdAlternateEmail } from "react-icons/md";
import { PiPasswordFill } from "react-icons/pi";
import { Toaster, toast } from 'react-hot-toast';
import { FaEye, FaEyeSlash } from 'react-icons/fa';

function Login() {

    const navigate = useNavigate();


    const [email, setEmail] = useState('');
    const [contrasenia, setContrasenia] = useState('');
    const [error, setError] = useState('');
    const [mostrarContrasenia, setMostrarContrasenia] = useState(false);

    const handleLogin = async (e) => {
        e.preventDefault();
        setError('');
           

         try {
             const response = await axios.post('http://localhost:8080/login', {
                email: email,
                contrasenia: contrasenia

                
            });
            console.log(email,contrasenia)
            localStorage.setItem('token', response.data.token);
           
            toast.success('¡Ingreso exitoso!', {
                    duration: 3000, 
            });
            limpiarCampos()
            navigate("/")
            } catch (err) {
                setError('Correo o contraseña incorrectos');
                toast.error('Correo o contraseña incorrectos', {
                    duration: 3000, 
                });
            }
    };
    
    const limpiarCampos = () => {
        setEmail(''),
        setContrasenia('')
    }

    const visualizarContrasenia = (e) => {
        e.preventDefault();
        e.stopPropagation();
        setMostrarContrasenia(!mostrarContrasenia); 
    };

    return (
        <>
            <div className="loginWrapper">
                <div className="loginContainer">
                    <div className="infoLogin">
                        <form>
                            <h1>INICIAR SESIÓN</h1>
                            <div className='inputs-form'>
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
                                    <input 
                                        type={mostrarContrasenia ? 'text' : 'password'}
                                        placeholder='Contraseña'
                                        value={contrasenia}
                                        onChange={(e) => setContrasenia(e.target.value)}
                                        required
                                    />
                                    <PiPasswordFill className='icon2' />
                                    <button 
                                        type="button"
                                        className="toggle-password"
                                        onClick={visualizarContrasenia}
                                        style={{ background: 'none', border: 'none', padding: 0 }}
                                    >
                                        {mostrarContrasenia ? (
                                            <FaEyeSlash className='icon-contrasenia' />
                                        ) : (
                                            <FaEye className='icon-contrasenia' />
                                        )}
                                    </button>
                                </div>
                            </div>
                            <div className="forgot-password">
                                <a href="#">¿Olvidaste la Contraseña?</a>
                            </div>
                            <button className="main-button" onClick={handleLogin}>Ingresar</button>

                            <div className="new-account">
                                <p>¿Eres nuevo aquí? <a href="/Registro">Crear cuenta</a></p>
                            </div>
                        </form>
                    </div>
                    <div className="imgLogin">
                        <img src={imagenLogin} alt="Imagen Gimnasio" />
                    </div>
                </div>
            </div>
        </>
    );
        
    

}

export default Login