import { Bars3Icon } from '../../ui/icons'

interface AdminHeaderProps {
  onMenuClick: () => void
}

const AdminHeader = ({ onMenuClick }: AdminHeaderProps) => {
  return (
    <header className="bg-white shadow-sm z-20">
      <div className="flex items-center justify-between px-6 py-4">
        <div className="flex items-center">
          <button 
            onClick={onMenuClick}
            className="mr-4 text-gray-600 hover:text-purple-600 md:hidden"
          >
            <Bars3Icon className="w-6 h-6" />
          </button>
          <h2 className="text-xl font-semibold text-gray-800">Admin Panel</h2>
        </div>
        
        <div className="flex items-center space-x-4">
          <div className="relative">
            <button className="flex items-center space-x-2">
              <img 
                src="https://placehold.co/400" 
                alt="User" 
                className="w-8 h-8 rounded-full"
              />
              <span className="hidden md:inline text-gray-700">Admin</span>
            </button>
          </div>
        </div>
      </div>
    </header>
  )
}

export default AdminHeader