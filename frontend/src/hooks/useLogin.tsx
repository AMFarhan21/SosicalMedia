import { useState } from 'react'
import { useNavigate } from 'react-router-dom'


const useLogin = () => {
    const [error, setError] = useState<string | null>(null)
    const [loading, setLoading] = useState(false)
    const navigate = useNavigate()

    const HOST = import.meta.env.VITE_API_HOST

    const login = async (email: string, password: string) => {
        try {
            setLoading(true)
            const res = await fetch(`${HOST}/api/v1/auth/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ email, password })
            })

            const result = await res.json()

            if (!res.ok) {
                throw new Error(result.error || "Login failed")
            }

            sessionStorage.setItem("Token", result.data)
            navigate("/", { replace: true })

        } catch (err) {
            if (err instanceof Error) {
                setError(err.message)
            } else {
                setError("Login failed")
            }
        } finally {
            setLoading(false)
        }
    }


    return { login, error, loading }
}

export default useLogin