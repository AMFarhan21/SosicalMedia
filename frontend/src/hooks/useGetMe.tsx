import { useEffect, useState } from "react"

interface userProperties {
    "id": string,
    "first_name": string,
    "last_name": string,
    "address": string,
    "email": string,
    "username": string,
    "age": number,
    "role": string,
}

const useGetMe = () => {
    const HOST = import.meta.env.VITE_API_HOST
    const token = localStorage.getItem("Token")
    const [Me, setMe] = useState<userProperties>({
        "id": "",
        "first_name": "",
        "last_name": "",
        "address": "",
        "email": "",
        "username": "",
        "age": 0,
        "role": "",
    })
    const [errorGetMe, setError] = useState("")
    const [loadingGetMe, setLoading] = useState(false)

    useEffect(() => {
        const getMe = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${HOST}/api/v1/users`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })
                const result = await res.json()
                if (!res.ok) {
                    if (res.status == 401 || res.status == 403) {
                        throw new Error("UNAUTHORIZED")
                    }

                    throw new Error(result.err)
                }

                setMe(result.data)
            } catch (err) {
                console.error(err)
                if (err instanceof Error) {
                    setError(err.message)
                }
            } finally {
                setLoading(false)
            }
        }

        getMe()
    }, [HOST, token])

    return { Me, errorGetMe, loadingGetMe }
}

export default useGetMe