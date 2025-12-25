import { useState } from 'react'
import type { PostsWithUsername } from './useGetAllPosts'

const useDeletePosts = (setPosts?: React.Dispatch<React.SetStateAction<PostsWithUsername[]>>) => {
    const [errorDelete, setError] = useState("")

    const HOST = import.meta.env.VITE_API_HOST
    const token = localStorage.getItem("Token")

    const deletePost = async (id: number) => {
        try {
            const res = await fetch(`${HOST}/api/v1/posts/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            const result = await res.json()

            if (!res.ok) {
                throw new Error(result.error)
            }

            if (setPosts) {
                setPosts(prev => [...prev].filter((p) => p.id != id))
            }
            setError("")
            return result.success
        } catch (err) {
            if (err instanceof Error) {
                setError("Failed to delete post, please try again!")
            } else {
                setError("Error on delete post")
            }
        }
    }


    return { deletePost, errorDelete }
}

export default useDeletePosts