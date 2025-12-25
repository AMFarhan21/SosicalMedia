import type { CommentsWithUsername } from "./useGetPostIDComments"

const useDeleteComments = () => {
    const deleteComments = async (post_id: number, comment_id: number, setComments: React.Dispatch<React.SetStateAction<CommentsWithUsername[]>>) => {
        const token = localStorage.getItem("Token")
        const HOST = import.meta.env.VITE_API_HOST
        try {
            const res = await fetch(`${HOST}/api/v1/posts/${post_id}/comments/${comment_id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })

            const result = await res.json()

            if (!res.ok) {
                throw new Error(result.message)
            }

            setComments((prev) => [...prev].filter((comment) => comment.id != comment_id))

        } catch (err) {
            console.error(err)
        }
    }

    return { deleteComments }
}

export default useDeleteComments