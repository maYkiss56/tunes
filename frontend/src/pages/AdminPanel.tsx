import { useEffect, useRef, useState } from 'react'
import { Outlet } from 'react-router-dom'
import AdminHeader from '../components/blocks/admin/AdminHeader'
import AdminSidebar from '../components/blocks/admin/AdminSidebar'

const AdminPanel = () => {
	const [sidebarOpen, setSidebarOpen] = useState(false)
	const sidebarRef = useRef<HTMLDivElement>(null)

	useEffect(() => {
		const handleClickOutside = (event: MouseEvent) => {
			if (
				sidebarOpen &&
				sidebarRef.current &&
				!sidebarRef.current.contains(event.target as Node)
			) {
				setSidebarOpen(false)
			}
		}

		document.addEventListener('mousedown', handleClickOutside)
		return () => {
			document.removeEventListener('mousedown', handleClickOutside)
		}
	}, [sidebarOpen])

	return (
		<div className='flex h-screen bg-gray-100 overflow-hidden'>
			<div
				ref={sidebarRef}
				className={`fixed inset-y-0 left-0 z-30 w-64 transform ${
					sidebarOpen ? 'translate-x-0' : '-translate-x-full'
				} md:translate-x-0 transition-transform duration-200 ease-in-out`}
			>
				<AdminSidebar onClose={() => setSidebarOpen(false)} />
			</div>

			<div className='flex-1 flex flex-col overflow-hidden md:ml-64'>
				<AdminHeader onMenuClick={() => setSidebarOpen(!sidebarOpen)} />

				<main className='flex-1 overflow-y-auto p-6 bg-gradient-to-b from-purple-50 to-gray-100'>
					<div className='max-w-7xl mx-auto'>
						<Outlet />
					</div>
				</main>
			</div>
		</div>
	)
}

export default AdminPanel
