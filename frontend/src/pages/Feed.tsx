import { useState } from 'react'
import useCreatePosts from '../hooks/useCreatePosts'
import Posts from '../components/Posts'
import useGetAllPosts from '../hooks/useGetAllPosts'
import { House, Search, User } from 'lucide-react'

const Feed = () => {
    const [content, setContent] = useState("")
    const [imageUrl, setImageUrl] = useState("")

    const { posts, setPosts } = useGetAllPosts()
    const { createPost, errorCreate } = useCreatePosts(setPosts)

    const handleCreatePost = (e: React.FormEvent) => {
        e.preventDefault()
        createPost(content, imageUrl)


        setContent("")
        setImageUrl("")
    }

    return (
        <div>
            <div className='w-full sm:w-[40%]'>
                <House />
                <Search />
                <User />

            </div>
            <div className='w-full sm:w-[40%] mx-auto'>
                <div className='font-bold text-2xl text-center mx-auto mb-6'> Feed </div>
                <form onSubmit={handleCreatePost} className='border border-gray-400 rounded-xl relative cursor-pointer p-4 flex flex-col space-y-4 mb-2'>
                    <input value={content} onChange={(e) => setContent(e.target.value)} placeholder='Create....' className='border-none outline-none' />
                    <input value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} placeholder='image_url' className='border-none outline-none text-sm' />
                    <button type='submit' className='bg-pink-800 cursor-pointer text-white w-20 rounded-lg absolute right-4 bottom-0 pb-1'>
                        create
                    </button>
                    <div className='text-sm text-red-500'>
                        {errorCreate}
                    </div>
                </form>
                <Posts posts={posts} setPosts={setPosts} />
            </div>
        </div>
    )

}

export default Feed