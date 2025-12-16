import { useState } from "react"
import useLogin from "../hooks/useLogin"

const Login = ({ setIsLogin }: { setIsLogin: React.Dispatch<React.SetStateAction<boolean>> }) => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const { login, loading, error } = useLogin()


    const handleLoginSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        login(email, password)
    }


    return (
        <div className="flex flex-col m-auto">
            <div className="ml-8">
                <div className="text-7xl font-bold mb-10">
                    Happening now
                </div>
                <div className="text-4xl font-bold mb-10">
                    Join today.
                </div>
            </div>
            <div className="w-lg m-auto border border-white/25 rounded-2xl p-4">
                <div className="text-center mb-4 font-bold text-2xl">Login</div>
                <form className="space-y-4" onSubmit={handleLoginSubmit}>
                    <input placeholder="Email" type="email" onChange={(e) => setEmail(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                    <input placeholder="Password" type="password" onChange={(e) => setPassword(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                    {
                        error && (
                            <div className="text-red-500 text-sm text-center">
                                {error}
                            </div>
                        )
                    }
                    <button disabled={loading} type="submit" className="mx-auto flex bg-white text-black px-8 py-1 rounded-lg cursor-pointer hover:bg-gray-200 font-bold">
                        {
                            loading ? "Loading..." : "Login"
                        }
                    </button>
                </form>
                <div className="text-gray-500 mt-4">
                    Don't have an account? <button type="button" onClick={() => setIsLogin(false)} className="text-blue-500 cursor-pointer">Sign up</button>
                </div>
            </div>
        </div>
    )
}

export default Login