import logoJartz from '../images/logoJartzCompletoBlanco.png'
import { FaFacebook } from "react-icons/fa";
import { BsInstagram } from "react-icons/bs";
import '../styles/footer.css'
function Footer() {

    const nombreEmpresa = "Gimnasio"
    return (
        <>
        <div className="containerFooter">
            <div className='line'>-</div>
            <div className="logoFooter">
                <img src={logoJartz} />
            </div>
            <div className='socialNetworks'>
                <FaFacebook onClick={() => window.open("https://www.facebook.com")} />

                <BsInstagram onClick={() => window.open("https://www.instagram.com")} />
            </div>
            <div className='infoFooter'>
                <p>&copy; 2025 {nombreEmpresa}. Todos los derechos reservados.</p>
            </div>
        </div>
        </>

    )
}
export default Footer