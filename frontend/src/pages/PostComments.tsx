import { useParams } from "react-router-dom"
import useGetPostIDComments from "../hooks/useGetPostIDComments"
import { Ellipsis, SquarePen, Trash } from 'lucide-react'
import { useState } from 'react'
import defaultProfile from '../assets/defaultProfile.jpg'
import { IoChatbubbleOutline, IoHeartOutline, IoHeartSharp } from 'react-icons/io5'
import Post from "../components/Post"



const PostComments = () => {
    const param = useParams()
    const [isLike, setIsLike] = useState(false)
    const [isComment, setIsComment] = useState(false)
    const [open, setOpen] = useState(false)

    const { post, comments, error } = useGetPostIDComments(Number(param.postID))
    if (error != "") {
        alert(error)
    }


    return (
        <div className="w-full sm:w-[56%] mx-auto">
            {
                post && <Post post={post} setPosts={undefined} idxCondition={true} postCommentsPage={() => { }} onPostComment={true} />
            }
            <div>
                {
                    comments.map((comment) => (
                        <div className={`border-t p-4 sm:border-l sm:border-r border-b border-white/22 cursor-pointer relative`} key={comment?.id}>
                            <div className='flex justify-between'>
                                <div className='flex gap-2'>
                                    <img src={defaultProfile} alt="default" className='w-8 rounded-full' />
                                    <div className='flex gap-2 items-center'>
                                        <div className='font-bold'>
                                            {comment?.first_name + " " + comment?.last_name}
                                        </div>
                                        <div className='text-gray-400'>
                                            @{comment?.username}
                                        </div>
                                        <div className='text-gray-400 text-xs'>
                                            {comment?.updated_at.split("T")[0]}
                                        </div>

                                    </div>
                                </div>
                                <button onClick={() => setOpen(true)} className='cursor-pointer hover:bg-gray-200 px-1 rounded-lg'>
                                    <Ellipsis />
                                </button>
                            </div>

                            <div className='my-4'>
                                {comment?.content}
                            </div>
                            {
                                comment?.image_url && <img className='mb-2' src={comment?.image_url} />
                            }

                            <div className='flex justify-around space-x-4'>
                                <button onClick={() => setIsComment(!isComment)} className='cursor-pointer gap-1'>
                                    <div className='flex gap-1 hover:text-blue-400 text-xs duration-200 items-center'>
                                        <div className='hover:bg-blue-300/20 rounded-full p-2 duration-200'>
                                            <IoChatbubbleOutline className='w-4 h-4' />
                                        </div>
                                        <span className='-ml-1'>
                                            0
                                        </span>
                                    </div>
                                </button>
                                <button onClick={() => setIsLike(!isLike)} className='cursor-pointer'>
                                    {
                                        isLike ? (
                                            <div className='flex gap-1 hover:text-pink-400 text-xs duration-200 items-center'>
                                                <div className='hover:bg-pink-300/20 rounded-full p-2 duration-200'>
                                                    <IoHeartOutline className='w-4 h-4' />
                                                </div>
                                                <span className='-ml-1'>
                                                    0
                                                </span>
                                            </div>
                                        ) : (
                                            <div className='flex gap-1 hover:text-pink-400 text-xs duration-200 items-center'>
                                                <div className='hover:bg-pink-300/20 rounded-full p-2 duration-200'>
                                                    <IoHeartSharp className='w-4 h-4 text-pink-600' />
                                                </div>
                                                <span className='-ml-1'>
                                                    0
                                                </span>
                                            </div>
                                        )

                                    }
                                </button>

                            </div>
                            {
                                open && (
                                    <div className='bg-white border p-2 rounded-lg absolute right-0 top-10 -mr-10'>
                                        <button onClick={async () => {
                                            // const success = awai(comment!.id)
                                            // if (!success) {
                                            //     alert(errorDelete)
                                            // } else {
                                            //     navigate(-1)
                                            // }

                                        }} className='flex gap-2 hover:bg-gray-200 px-2 rounded-sm w-full cursor-pointer text-red-500'>
                                            <Trash className='w-4' />
                                            <div>
                                                delete
                                            </div>
                                        </button>
                                        <button className='flex gap-2 hover:bg-gray-200 px-2 rounded-sm w-full cursor-pointer text-blue-400'>
                                            <SquarePen className='w-4' />
                                            <div>
                                                edit
                                            </div>
                                        </button>
                                    </div>
                                )
                            }
                        </div>
                    ))
                }
            </div>
        </div>
    )
}

export default PostComments