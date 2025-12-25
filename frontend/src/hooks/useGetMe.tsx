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

    useEffect(() => {
        const getMe = async () => {
            try {
                const res = await fetch(`${HOST}/api/v1/users`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })
                const result = await res.json()
                if (!res.ok) {
                    throw new Error(result.err)
                }

                setMe(result.data)
            } catch (err) {
                console.error(err)
            }
        }

        getMe()
    }, [HOST, token])

    return { Me }
}

export default useGetMe