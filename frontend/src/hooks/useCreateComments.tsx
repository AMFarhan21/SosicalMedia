import { useState } from "react"
import type { CommentsWithUsername } from "./useGetPostIDComments"

const useCreateComments = ({ setComments }: { setComments: React.Dispatch<React.SetStateAction<CommentsWithUsername[]>> }) => {
    const [errorComments, setError] = useState("")
    const HOST = import.meta.env.VITE_API_HOST
    const token = sessionStorage.getItem("Token")

    const createComments = async (post_id: number, content: string, files: File[]) => {
        try {

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

            setComments((prev) => [resultGetByID.data, ...prev])
        } catch (err) {
            if (err instanceof Error) {
                console.error(err.message)
                setError(errorComments)
            } else {
                console.error("Error on creating comment")
                setError("Error on creating comment")
            }
        }
    }

    return { createComments, errorComments }
}

export default useCreateComments