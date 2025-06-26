import { Link } from 'react-router-dom'
import { AlbumIcon, GenreIcon, HomeIcon, MusicalNoteIcon, UsersIcon } from '../../ui/icons'

interface AdminSidebarProps {
  onClose: () => void
}

const AdminSidebar = ({ onClose }: AdminSidebarProps) => {
  return (
    <aside className="w-64 h-full bg-gray-900 text-white shadow-lg">
      <div className="p-4 border-b border-gray-700">
        <h1 className="text-2xl font-bold text-purple-400">MelodyCritic Admin</h1>
      </div>
      
      <nav className="p-4">
        <ul className="space-y-2">
          <li>
            <Link 
              to="/admin" 
              className="flex items-center p-3 rounded-lg hover:bg-purple-800 transition-colors"
              onClick={onClose}
            >
              <HomeIcon className="w-5 h-5 mr-3" />
                Домой 
            </Link>
          </li>
          <li>
            <Link 
              to="/admin/songs" 
              className="flex items-center p-3 rounded-lg hover:bg-purple-800 transition-colors"
              onClick={onClose}
            >
              <MusicalNoteIcon className="w-5 h-5 mr-3" />
              Песни
            </Link>
          </li>
          <li>
            <Link 
              to="/admin/artists" 
              className="flex items-center p-3 rounded-lg hover:bg-purple-800 transition-colors"
              onClick={onClose}
            >
              <UsersIcon className="w-5 h-5 mr-3" />
              Артисты
            </Link>
          </li>
          <li>
            <Link 
              to="/admin/albums" 
              className="flex items-center p-3 rounded-lg hover:bg-purple-800 transition-colors"
              onClick={onClose}
            >
              <AlbumIcon className="w-5 h-5 mr-3" />
              Альбомы
            </Link>
          </li>
          <li>
            <Link 
              to="/admin/genres" 
              className="flex items-center p-3 rounded-lg hover:bg-purple-800 transition-colors"
              onClick={onClose}
            >
              <GenreIcon className="w-5 h-5 mr-3" />
              Жанры
            </Link>
          </li>
        </ul>
      </nav>
    </aside>
  )
}

export default AdminSidebar