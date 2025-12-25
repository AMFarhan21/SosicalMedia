import { useParams } from "react-router-dom"
import useGetPostIDComments from "../hooks/useGetPostIDComments"
import { Image, X } from 'lucide-react'
import { useEffect, useState } from 'react'
import Post from "../components/Post"
import Comments from "../components/Comments"
import useCreateComments from "../hooks/useCreateComments"
import LoadingSkeleton from "../components/LoadingSkeleton"



const PostComments = () => {
    const [commentInput, setCommentInput] = useState("")
    const [img, setImg] = useState("")
    const [showImg, setShowImg] = useState(false)
    const [preview, setPreview] = useState("")
    const [files, setFiles] = useState<File[]>([])

    const param = useParams()
    const HOST = import.meta.env.VITE_API_HOST

    const { post, setPost, comments, setComments, error, loadingPost, loadingComments } = useGetPostIDComments(Number(param.postID))


    const { createComments, errorComments, loadingCreateComment } = useCreateComments()


    const handleCreateComment = async (e: React.FormEvent) => {
        e.preventDefault()
        if (!post) return
        const res = await createComments(post.id, commentInput, files, setComments)
        if (res) {
            setCommentInput("")
            setFiles([])
        }
    }


    useEffect(() => {
        if (error != "") {
            alert(error)
        }
    }, [error])

    useEffect(() => {
        if (img || preview) {
            const timer = setTimeout(() => setShowImg(true), 10)
            return () => clearTimeout(timer)
        } else {
            const timer = setTimeout(() => setShowImg(false), 10)
            return () => clearTimeout(timer)
        }
    }, [img, preview])




    return (
        <>
            {
                img ? (
                    <div className={`fixed inset-0 bg-black/80 z-51 items-center m-auto flex transition-opacity scale-100 duration-300 ${showImg ? "opacity-100" : "opacity-0"}`}>
                        <img loading="lazy" src={`${HOST}/${img}`} className={`max-h-full mx-auto transition-opacity scale-100 duration-300 ${showImg ? "scale-100" : "scale-0"}`} />
                        <div onClick={() => setImg("")} className='absolute top-10 right-10 cursor-pointer'>
                            <X />
                        </div>
                    </div>
                ) : preview && (
                    <div className={`fixed inset-0 bg-black/80 z-51 items-center m-auto flex transition-opacity scale-100 duration-300 ${showImg ? "opacity-100" : "opacity-0"}`}>
                        <img loading='lazy' src={preview} className={`max-h-full mx-auto transition-opacity scale-100 duration-300 ${showImg ? "scale-100" : "scale-0"}`} />
                        <div onClick={() => setPreview("")} className='absolute top-10 right-10 cursor-pointer bg-black/50 rounded-full p-1 m-1 hover:bg-black/70'>
                            <X />
                        </div>
                    </div>
                )
            }
            <div className="w-full sm:w-[56%] mx-auto">
                {
                    loadingPost ? (
                        <div className="border border-white/22 p-2">
                            <LoadingSkeleton />
                        </div>
                    ) : (
                        <Post post={post} setPost={setPost} idxCondition={true} postCommentsPage={() => { }} onPostComment={true} setImg={setImg} />
                    )
                }
                <form onSubmit={handleCreateComment} className={`sm:border-x ${!comments && "border-b"} border-white/22 relative cursor-pointer p-4 flex flex-col space-y-4`}>
                    <textarea value={commentInput} onChange={(e) => setCommentInput(e.target.value)} placeholder={`Reply to @${post.username}`} className="outline-none" />
                    <div className='flex flex-wrap'>
                        {
                            files && files.map((file, index) => (
                                <div key={index} className='relative w-[50%]'>
                                    <button type='button' className='absolute right-0 bg-black/50 rounded-full p-1 m-1 cursor-pointer hover:bg-black/70' onClick={() => setFiles((prev) => prev.filter((_, i) => i != index))}>
                                        <X className='' />
                                    </button>
                                    <img onClick={() => setPreview(URL.createObjectURL(file))} src={URL.createObjectURL(file)} className='' />
                                </div>
                            ))
                        }
                    </div>
                    <label htmlFor="upload-image">
                        <div className="cursor-pointer hover:bg-blue-400/20 duration-100 p-2 w-8 rounded-full -mb-6">
                            <Image className="w-4 h-4 text-blue-500" />
                        </div>
                    </label>
                    <input id="upload-image" type="file" accept="image/*" multiple onChange={(e) => {
                        if (!e.target.files) return
                        setFiles(Array.from(e.target.files))
                    }} placeholder={`Image url`} className="hidden outline-none" />

                    <button disabled={loadingCreateComment} className={`cursor-pointer absolute right-4 bottom-4 bg-white ${loadingCreateComment ? "bg-white/50 hover:bg-white/50" : "hover:bg-white/80"} cursor-pointer text-black font-bold px-4 py-1 cursor-pointerduration-100 rounded-full`} type="submit">
                        Reply
                    </button>
                    <div className='text-sm text-red-500'>
                        {errorComments}
                    </div>
                </form>
                <div>
                    {
                        loadingComments ? (
                            Array.from({ length: post.comments_count }).map(() => (
                                <div className="border border-white/22 p-2">
                                    <LoadingSkeleton />
                                </div>
                            ))
                        ) : comments && comments.map((comment) => (
                            <Comments post_id={post.id} username={post.username} comment={comment} setComments={setComments} />
                        ))
                    }
                </div>
            </div>
        </>
    )
}

export default PostComments