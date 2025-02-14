import imagenLogin from '../images/imagenLogin.png'
import '../styles/login.css'
import {useNavigate } from "react-router-dom"
import { FaRegUser } from "react-icons/fa";
import { PiPasswordFill } from "react-icons/pi";
import { Toaster, toast } from 'react-hot-toast';

function Login() {

    const navigate = useNavigate();

    const irInicio = (e) => {
        e.preventDefault();
        toast.success('¡Ingreso exitoso!', {
            duration: 3000, 
        });
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
                                    <input type='text' placeholder='Usuario' />
                                    <FaRegUser className='icon' />
                                </div>
                                <div className='input-container'>
                                    <input type='password' placeholder='Contraseña' />
                                    <PiPasswordFill className='icon' />
                                </div>
                            </div>
                            <div className="forgot-password">
                                <a href="#">¿Olvidaste la Contraseña?</a>
                            </div>
                            <button onClick={irInicio}>Ingresar</button>

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