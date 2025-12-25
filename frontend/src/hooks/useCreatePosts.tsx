import { useState } from "react"
import type { PostsWithUsername } from "./useGetAllPosts"

interface CreatePostsResponse {
    data: PostsWithUsername
}

const useCreatePosts = (setPosts: React.Dispatch<React.SetStateAction<PostsWithUsername[]>>) => {
    const [errorCreate, setError] = useState("")
    const [loadingCreatePost, setLoading] = useState(false)

    const token = localStorage.getItem("Token")
    const HOST = import.meta.env.VITE_API_HOST

    const createPost = async (content: string, files: File[]) => {
        try {
            setLoading(true)
            const formData = new FormData()
            formData.append("content", content)
            for (const file of files) {
                formData.append("images", file)
            }
            const resCreate = await fetch(`${HOST}/api/v1/posts`, {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${token}`
                },
                body: formData
            })

            const resultCreate = await resCreate.json()

            if (!resCreate.ok) {
                throw new Error(resultCreate.error)
            }

            const resGetID = await fetch(`${HOST}/api/v1/posts/${resultCreate.data.id}`, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            const resultGetID: CreatePostsResponse = await resGetID.json()


            setPosts(prev => [resultGetID.data, ...prev])
            setError("")
            return true
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message)
            } else {
                setError("Error on creating post")
            }
        } finally {
            setLoading(false)
        }
    }

    return { createPost, errorCreate, loadingCreatePost }
}

export default useCreatePosts