import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './styles/index.css'
import { GoogleOAuthProvider } from '@react-oauth/google'
import { NOTELAD_CLIENT_ID } from './consts/endpoints.ts'
import { BrowserRouter } from 'react-router-dom'




ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId={NOTELAD_CLIENT_ID}>
    <BrowserRouter>
    <App />
    </BrowserRouter>
    </GoogleOAuthProvider>
  </React.StrictMode>,
)
