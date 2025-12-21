import { useState } from "react"
import xlogo from "../assets/x2.png"
import Register from "../components/Register"
import Login from "../components/Login"
const Auth = () => {

    const [isLogin, setIsLogin] = useState(true)

    return (
        <div className="gap-20 w-full min-h-screen relative grid md:grid-cols-2 items-center justify-center m-auto flex-wrap">
            <img src={xlogo} alt="App Logo" className='w-[60%] text-center items-center mx-auto' />
            <div className="">
                <div className="text-4xl sm:text-7xl font-bold mb-10">
                    Happening now
                </div>
                <div className="text-2xl sm:text-4xl font-bold mb-10">
                    Join today.
                </div>
                {
                    isLogin ? (
                        <Login setIsLogin={setIsLogin} />
                    ) : (
                        <Register setIsLogin={setIsLogin} />
                    )
                }
            </div>
        </div>
    )
}

export default Auth