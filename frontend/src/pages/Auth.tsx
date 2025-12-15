import { useState } from "react"
import useLogin from "../hooks/useLogin"

const Auth = () => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const { login, loading, error } = useLogin()

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        login(email, password)
    }


    return (
        <div className="w-full h-screen -mt-20 flex">
            <div className="w-lg m-auto border rounded-2xl p-4 justify-center items-center">
                <div className="text-center mb-4 font-bold text-2xl">Login</div>
                <form className="space-y-2" onSubmit={handleSubmit}>
                    <input placeholder="Email" type="email" onChange={(e) => setEmail(e.target.value)} className="w-full border rounded-lg px-4 py-2" />
                    <input placeholder="Password" type="password" onChange={(e) => setPassword(e.target.value)} className="w-full border rounded-lg px-4 py-2" />
                    {
                        error && (
                            <div className="text-red-500 text-sm text-center">
                                {error}
                            </div>
                        )
                    }
                    <button disabled={loading} type="submit" className="mx-auto flex bg-black text-white px-8 py-1 rounded-lg cursor-pointer hover:bg-gray-600">
                        {
                            loading ? "Loading..." : "Login"
                        }
                    </button>
                </form>
            </div>
        </div>
    )
}

export default Auth