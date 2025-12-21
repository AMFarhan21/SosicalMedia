import React, { useState } from 'react'
import useRegister from '../hooks/useRegister'

const Register = ({ setIsLogin }: { setIsLogin: React.Dispatch<React.SetStateAction<boolean>> }) => {
    const [firstName, setFirstName] = useState("")
    const [lastName, setLastName] = useState("")
    const [address, setAddress] = useState("")
    const [email, setEmail] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [age, setAge] = useState(0)
    const { register, registerLoading, registerError } = useRegister()

    const handleRegisterSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        register(firstName, lastName, address, email, username, password, age)
    }
    return (
        <div className="w-full m-auto border border-white/25 rounded-2xl p-4">
            <div className="text-center mb-4 font-bold text-2xl">Register</div>
            <form className="space-y-4" onSubmit={handleRegisterSubmit}>
                <div className='flex gap-2'>
                    <input placeholder="First name" onChange={(e) => setFirstName(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                    <input placeholder="Last name" onChange={(e) => setLastName(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                </div>
                <input placeholder="Address" onChange={(e) => setAddress(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                <input placeholder="Email" type="email" onChange={(e) => setEmail(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                <input placeholder="Username" onChange={(e) => setUsername(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                <input placeholder="Password" type="password" onChange={(e) => setPassword(e.target.value)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                <input placeholder="Age" type='number' onChange={(e) => setAge(e.target.valueAsNumber)} className="w-full border border-white/25 rounded-lg px-4 py-2" />
                {
                    registerError && (
                        <div className="text-red-500 text-sm text-center">
                            {registerError}
                        </div>
                    )
                }
                <button disabled={registerLoading} type="submit" className="bg-white text-black w-full py-1 rounded-lg cursor-pointer hover:bg-gray-200 font-bold">
                    {
                        registerLoading ? "Loading..." : "Register"
                    }
                </button>

            </form>
            <div className="text-gray-500 mt-4">
                Don't have an account? <button type="button" onClick={() => setIsLogin(true)} className="text-blue-500 cursor-pointer">Sign in</button>
            </div>
        </div>
    )
}

export default Register