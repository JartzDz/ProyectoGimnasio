import Register from '../components/Register'
import Footer from '../components/Footer'
import '../styles/register.css'
import {Toaster} from 'react-hot-toast'

function RegisterPage () {
    return (
        <>
         <div className="pageWrapper">
            <Register/>
            <Footer/>
            <div><Toaster
                position="top-right"
                reverseOrder={false}
            />
            </div>
        </div>
        </>

    )
}
export default RegisterPage