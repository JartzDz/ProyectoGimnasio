import Login from '../components/Login';
import Footer from '../components/Footer';
import '../styles/login.css'
import {Toaster} from 'react-hot-toast'
function LoginPage() {
    return (
        <div className="pageWrapper">
            <Login />
            <Footer />
            <div><Toaster
                position="top-right"
                reverseOrder={false}
            />
            </div>
        </div>
    );
}

export default LoginPage;
