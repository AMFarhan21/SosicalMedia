import { useState } from "react"
import xlogo from "../assets/x2.png"
import Register from "../components/Register"
import Login from "../components/Login"
const Auth = () => {

    const [isLogin, setIsLogin] = useState(true)

    return (
        <div className="w-full h-full left-10 flex absolute">
            <div className="m-auto -mr-20">
                <img src={xlogo} alt="App Logo" className='w-68' />
            </div>
            {
                isLogin ? (
                    <Login setIsLogin={setIsLogin} />
                ) : (
                    <Register setIsLogin={setIsLogin} />
                )
            }
        </div>
    )
}

export default Auth