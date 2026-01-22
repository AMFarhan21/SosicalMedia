import { useState } from "react"
import { useNavigate } from "react-router-dom"
import useLogin from "./useLogin"

const useRegister = () => {
    const [registerError, setError] = useState("")
    const [registerLoading, setLoading] = useState(false)

    const HOST = import.meta.env.VITE_API_HOST

    const { login } = useLogin()

    const navigate = useNavigate()

    const register = async (firstName: string, lastName: string, email: string, username: string, password: string) => {
        try {
            setLoading(true)
            const res = await fetch(`${HOST}/api/v1/auth/register`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ first_name: firstName, last_name: lastName, email, username, password })
            })

            const result = await res.json()

            if (!res.ok) {
                throw new Error(result.error)
            }

            login(email, password)

            navigate("/", { replace: true })
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message)
            } else {
                setError("Failed to register")
            }
        } finally {
            setLoading(false)
        }
    }

    return { register, registerLoading, registerError }

}

export default useRegister