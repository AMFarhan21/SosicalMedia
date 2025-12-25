import { useState } from "react"
import type { CommentsWithUsername } from "./useGetPostIDComments"

const useCreateComments = () => {
    const [errorComments, setError] = useState("")
    const [loadingCreateComment, setLoading] = useState(false)
    const HOST = import.meta.env.VITE_API_HOST
    const token = localStorage.getItem("Token")

    const createComments = async (post_id: number, content: string, files: File[], setComments: React.Dispatch<React.SetStateAction<CommentsWithUsername[]>>) => {
        try {
            setLoading(true)
            const formData = new FormData()
            formData.append("content", content)
            for (const file of files) {
                formData.append("images", file)
            }

            const resCreate = await fetch(`${HOST}/api/v1/posts/${post_id}/comments`, {
                method: "POST",
                headers: {
                    "Authorization": `Bearer ${token}`
                },
                body: formData
            })
            const resultCreate = await resCreate.json()
            if (!resCreate.ok) {
                throw new Error(resultCreate.error)
            }

            const resGetByID = await fetch(`${HOST}/api/v1/posts/${post_id}/comments/${resultCreate.data.id}`, {
                headers: {
                    "Authorization": `Bearer ${token}`
                }
            })
            const resultGetByID = await resGetByID.json()
            if (!resGetByID.ok) {
                throw new Error(resultGetByID.error)
            }

            setComments((prev) => {
                if (!prev) {
                    return [resultGetByID.data]
                }
                return [resultGetByID.data, ...prev]
            })

            return true
        } catch (err) {
            if (err instanceof Error) {
                console.error(err.message)
                setError(err.message)
            } else {
                console.error("Error on creating comment")
                setError("Error on creating comment")
            }
        } finally {
            setLoading(false)
        }
    }

    return { createComments, errorComments, loadingCreateComment }
}

export default useCreateComments