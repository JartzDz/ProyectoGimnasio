import imagenLogin from '../images/imagenLogin.png'
import '../styles/login.css'
import React, { useState } from 'react';
import axios from 'axios';

import {useNavigate } from "react-router-dom"
import { MdAlternateEmail } from "react-icons/md";
import { PiPasswordFill } from "react-icons/pi";
import { Toaster, toast } from 'react-hot-toast';

function Login() {

    const navigate = useNavigate();


    const [email, setEmail] = useState('');
    const [contrasenia, setContrasenia] = useState('');
    const [error, setError] = useState('');

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
            navigate("/")
            } catch (err) {
                setError('Correo o contraseña incorrectos');
                toast.error('Correo o contraseña incorrectos', {
                    duration: 3000, 
                });
            }
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
                                        type='password' 
                                        placeholder='Contraseña'
                                        value={contrasenia}
                                        onChange={(e) => setContrasenia(e.target.value)}
                                        required
                                        />
                                    <PiPasswordFill className='icon' />
                                </div>
                            </div>
                            <div className="forgot-password">
                                <a href="#">¿Olvidaste la Contraseña?</a>
                            </div>
                            <button onClick={handleLogin}>Ingresar</button>

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
    )
        
    

}

export default Login