import { Ellipsis, Trash } from 'lucide-react'

import defaultProfile from '../assets/defaultProfile.jpg'

import { IoChatbubbleOutline, IoHeartOutline, IoHeartSharp } from 'react-icons/io5'
import React, { useEffect, useRef, useState } from 'react'
import type { CommentsWithUsername } from '../hooks/useGetPostIDComments'
import useLikes from '../hooks/useLikes'
import useDeleteComments from '../hooks/useDeleteComments'
import useGetMe from '../hooks/useGetMe'

const Comments = ({ post_id, username, comment, setComments }: { post_id: number, username: string, comment: CommentsWithUsername, setComments: React.Dispatch<React.SetStateAction<CommentsWithUsername[]>> }) => {

    const [isLiked, setIsLike] = useState(comment.is_liked)
    const [likesCount, setLikesCount] = useState(comment.likes_count)
    const [isComment, setIsComment] = useState(false)
    const [open, setOpen] = useState(false)

    const HOST = import.meta.env.VITE_API_HOST

    const { likes } = useLikes({ setComments })
    const { deleteComments } = useDeleteComments()
    const { Me } = useGetMe()

    const menuRef = useRef<HTMLDivElement | null>(null)
    useEffect(() => {
        const handleClickOutside = (e: MouseEvent) => {
            if (menuRef.current && !menuRef.current.contains(e.target as Node)) {
                setOpen(false)
            }
        }

        document.addEventListener("mousedown", handleClickOutside)
        return () => {
            document.removeEventListener("mousedown", handleClickOutside)
        }
    }, [])

    return (
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
                <button onClick={() => setOpen(true)} className='cursor-pointer hover:bg-white/10 px-1 rounded-lg'>
                    <Ellipsis />
                </button>
            </div>

            <div className='my-4'>
                <span className='text-blue-500'>@{username + " "}</span>
                {comment?.content}
            </div>
            <div className={`grid ${comment.image_url && comment.image_url.length > 1 ? "grid-cols-2" : "grid-cols-1"} gap-1`}>
                {
                    comment.image_url && comment.image_url.map((image, index) => (
                        <img className={`
                            ${comment.image_url.length == 4 && index == 0 ? "rounded-tl-xl" : comment.image_url.length == 4 && index == 1 ? "rounded-tr-xl" : comment.image_url.length == 4 && index == 2 ? "rounded-bl-xl" : comment.image_url.length == 4 && index == 3 && "rounded-br-xl"}
                            ${comment.image_url.length == 3 && index == 0 ? "row-span-2 h-full sm:h-full rounded-l-xl" : comment.image_url.length == 3 && index == 0 ? "rounded-tr-xl" : comment.image_url.length == 3 && index == 2 && "rounded-br-xl"}
                            ${comment.image_url.length == 1 ? "rounded-xl h-full w-full" : "h-40 sm:h-60 w-full"}
                            object-cover
                        `} src={`${HOST}/${image}`} />
                    ))

                }
            </div>
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
                <button onClick={() => setIsLike(!isLiked)} className='cursor-pointer'>
                    <div className='flex gap-1 hover:text-pink-400 text-xs duration-200 items-center'>
                        <button onClick={() => {
                            likes(comment.id, "COMMENT")
                            setLikesCount((prev) => {
                                if (isLiked) {
                                    return prev - 1
                                } else {
                                    return prev + 1
                                }
                            })
                        }} className='hover:bg-pink-300/20 rounded-full p-2 duration-200 cursor-pointer'>
                            {isLiked ? <IoHeartSharp className='w-4 h-4 text-pink-600' /> : <IoHeartOutline className='w-4 h-4' />}
                        </button>
                        <span className='-ml-1'>
                            {likesCount}
                        </span>
                    </div>
                </button>

            </div>
            {
                open && (
                    <div ref={menuRef} className='bg-black border border-white/20 p-1 rounded-lg absolute right-0 top-10'>
                        {
                            Me.id == comment.user_id ? (
                                <>
                                    <button onClick={async (e) => {
                                        e.preventDefault()
                                        deleteComments(post_id, comment.id, setComments)

                                    }} className='flex gap-2 hover:bg-white/10 px-2 rounded-sm w-full cursor-pointer text-red-500'>
                                        <Trash className='w-4' />
                                        <div>
                                            delete
                                        </div>
                                    </button>

                                </>
                            ) : (
                                <>
                                    <button className='flex gap-2 hover:bg-white/10 px-2 rounded-sm w-full cursor-pointer text-blue-400'>
                                        <div>
                                            (Work in Progress)
                                        </div>
                                    </button>
                                </>
                            )
                        }
                    </div>
                )
            }
        </div>
    )
}

export default Comments