import type { JSX } from "react"
import { Navigate } from "react-router-dom"

const ProtectedRoutes = ({ children }: { children: JSX.Element }) => {

    const token = sessionStorage.getItem("Token")

    if (!token) {
        return <Navigate to="/auth" replace />
    }


    return children
}

export default ProtectedRoutes