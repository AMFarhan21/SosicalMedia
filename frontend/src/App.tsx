import './App.css'
import { Link, Route, Routes, useNavigate } from 'react-router-dom'
import Feed from './pages/Feed'
import Auth from './pages/Auth'
import ProtectedRoutes from './components/ProtectedRoutes'

function App() {

  const navigate = useNavigate()
  const isLogin = sessionStorage.getItem("Token")

  return (
    <>
      <div className='flex flex-wrap justify-between mx-12 my-4 font-bold text-lg gap-2'>
        <div className='text-pink-800'>
          <Link to={"/"}> <img src='../assets/x.avif' /> </Link>
        </div>
        <div className='space-x-4'>
          <Link to={"/"} className='hover:text-gray-500'>Feed</Link>
          <Link to={"/profile"} className='hover:text-gray-500'>Profile</Link>
          {
            isLogin != "" ? (
              <button onClick={() => {
                sessionStorage.removeItem("Token")
                navigate("/auth")
              }} className='hover:text-gray-500 cursor-pointer'>Logout</button>
            ) : (
              <Link to={"/auth"} className='hover:text-gray-500'>Login</Link>
            )
          }

        </div>
      </div>

      <Routes>
        <Route path="/" element={
          <ProtectedRoutes>
            <Feed />
          </ProtectedRoutes>
        }></Route>
        <Route path="/auth" element={<Auth />}></Route>
      </Routes >
    </>
  )
}

export default App
