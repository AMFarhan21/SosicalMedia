import { useEffect, useState } from 'react'
import type { PostsWithUsername } from './useGetAllPosts'

export interface CommentsWithUsername {
    id: number
    user_id: string
    first_name: string
    last_name: string
    username: string
    content: string
    image_url: string
    created_at: string
    updated_at: string

}

const useGetPostIDComments = (postID: number) => {
    const [error, setError] = useState("")
    const [comments, setComments] = useState<CommentsWithUsername[]>([])
    const [post, setPost] = useState<PostsWithUsername | null>(null)
    const token = sessionStorage.getItem("Token")



    useEffect(() => {
        const fetchPost = async () => {
            try {
                const res = await fetch(`http://localhost:8000/api/v1/posts/${postID}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const result = await res.json()

                if (!res.ok) {
                    throw new Error(result.message)
                }

                setPost(result.data)
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message)
                } else {
                    setError("Error on fetching post by id")
                }
            }
        }
        const fetchComments = async () => {
            try {
                const res = await fetch(`http://localhost:8000/api/v1/posts/${postID}/comments`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })
                const result = await res.json()

                if (!res.ok) {
                    throw new Error(result.message)
                }


                setComments(result.data)
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message)
                } else {
                    setError("Error on fetching comments")
                }
            }
        }

        fetchPost()
        fetchComments()
    }, [postID, token])




    return { post, comments, error }

}

export default useGetPostIDComments