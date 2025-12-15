import { useState } from "react"
import type { PostsWithUsername } from "./useGetAllPosts"

interface CreatePostsResponse {
    data: PostsWithUsername
}

const useCreatePosts = (setPosts: React.Dispatch<React.SetStateAction<PostsWithUsername[]>>) => {
    const [errorCreate, setError] = useState("")

    const token = sessionStorage.getItem("Token")

    const createPost = async (content: string, image_url: string) => {
        try {
            const resCreate = await fetch("http://localhost:8000/api/v1/posts", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`
                },
                body: JSON.stringify({ content, image_url })
            })

            const resultCreate = await resCreate.json()

            if (!resCreate.ok) {
                throw new Error("Content should not be empty")
            }

            const resGetID = await fetch(`http://localhost:8000/api/v1/posts/${resultCreate.data.id}`, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            const resultGetID: CreatePostsResponse = await resGetID.json()


            setPosts(prev => [resultGetID.data, ...prev])
            setError("")
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message)
            } else {
                setError("Error on creating post")
            }
        }
    }

    return { createPost, errorCreate }
}

export default useCreatePosts