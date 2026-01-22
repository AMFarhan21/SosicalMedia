import './App.css'
import { Link, Route, Routes, useNavigate } from 'react-router-dom'
import Feed from './pages/Feed'
import Auth from './pages/Auth'
import ProtectedRoutes from './components/ProtectedRoutes'
import { House, LogIn, LogOut, Search, User } from 'lucide-react'
import xlogo from './assets/x2.png'
import PostComments from './pages/PostComments'
import useGetMe from './hooks/useGetMe'
import Profile from './pages/Profile'

function App() {

  const navigate = useNavigate()
  const isLogin = localStorage.getItem("Token")

  const { Me } = useGetMe()


  return (
    <div className='text-white bg-black min-h-screen'>

      <div className='w-full sm:w-[75%] mx-auto grid-cols-3'>
        {
          isLogin && (
            <>
              <div className='hidden sm:flex fixed space-y-8 flex-col'>
                <Link to={"/"} className='w-15 hover:bg-white/12 duration-100 p-4 rounded-full ml-1'> <img src={xlogo} alt="App Logo" className='' /> </Link>
                <Link to={"/"} className='flex gap-2 text-xl  items-center hover:bg-white/12 duration-100 p-2 rounded-full'> <House className='w-12' /> <span className='sm:hidden md:inline'>Home</span> </Link>
                <Link to={"/"} className='flex gap-2 text-xl  items-center hover:bg-white/12 duration-100 p-2 rounded-full'> <Search className='w-12' /> <span className='sm:hidden md:inline'>Search</span> </Link>
                <Link to={`/${Me.username}/${Me.id}`} className='flex gap-2 text-xl  items-center hover:bg-white/12 duration-100 p-2 rounded-full'> <User className='w-12' /> <span className='sm:hidden md:inline'>Profile</span> </Link>
                <div className='space-x-4 bottom-0'>
                  {
                    isLogin != "" ? (
                      <button onClick={() => {
                        localStorage.removeItem("Token")
                        navigate("/auth")
                      }} className='cursor-pointer flex gap-2 text-xl  hover:bg-white/12 duration-100 p-2 rounded-full'> <LogOut className='w-12' /> <span className='sm:hidden md:inline'>Logout</span> </button>
                    ) : (
                      <Link to={"/auth"} className='cursor-pointer flex gap-2 text-xl  hover:bg-white/12 duration-100 p-2 rounded-full'> <LogIn /> <span className='sm:hidden md:inline'>Login</span></Link>
                    )
                  }
                </div>
              </div>
              <div className='flex sm:hidden justify-around fixed bottom-0 bg-black w-full p-2 z-50'>
                <Link to={"/"} className='flex gap-2 text-xl font-bold items-center hover:bg-white/12 duration-100 p-2 rounded-full'> <House className='w-12' /> </Link>
                <Link to={"/"} className='flex gap-2 text-xl font-bold items-center hover:bg-white/12 duration-100 p-2 rounded-full'> <Search className='w-12' /> </Link>
                <Link to={`/${Me.username}/${Me.id}`} className='flex gap-2 text-xl font-bold items-center hover:bg-white/12 duration-100 p-2 rounded-full'> <User className='w-12' /> </Link>
                <div className='space-x-4'>
                  {
                    isLogin != "" ? (
                      <button onClick={() => {
                        localStorage.removeItem("Token")
                        navigate("/auth")
                      }} className='cursor-pointer flex gap-2 text-xl font-bold hover:bg-white/12 duration-100 p-2 rounded-full'> <LogOut className='w-12' /> </button>
                    ) : (
                      <Link to={"/auth"} className='cursor-pointer flex gap-2 text-xl font-bold hover:bg-white/12 duration-100 p-2 rounded-full'> <LogIn /> </Link>
                    )
                  }
                </div>
              </div>
            </>
          )
        }


        <Routes>
          <Route path="/" element={
            <ProtectedRoutes>
              <Feed />
            </ProtectedRoutes>
          }></Route>
          <Route path="/auth" element={<Auth />}></Route>
          <Route path="/:postID/comments" element={
            <ProtectedRoutes>
              <PostComments />
            </ProtectedRoutes>
          } />
          <Route path="/:username/:userID" element={
            <ProtectedRoutes>
              <Profile />
            </ProtectedRoutes>
          }></Route>
        </Routes >
      </div>
    </div>
  )
}

export default App
