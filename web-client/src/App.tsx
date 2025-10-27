import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Root from './Root.tsx';
import StartPage from './pages/StartPage.tsx'
import PoolPage from './pages/PoolPage.tsx'

function App() {

  return (
    <Routes>
        <Route path="/" element={<Root />}>
            <Route index element={<StartPage />} />
            <Route path="/pool" element={<PoolPage />} />
        </Route>
    </Routes>
  )
}

export default App
