import { type FC } from 'react'


const Footer: FC = () => {
  return (
    <footer className="bg-gradient-to-b from-gray-900 to-black text-gray-300 pt-16 pb-8 px-4 sm:px-6 lg:px-8 border-t border-gray-800">
      <div className="max-w-7xl mx-auto">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-12">
          
          {/* О приложении */}
          <div className="space-y-4">
            <h3 className="text-xl font-bold text-purple-400 flex items-center">
              <span className="text-2xl mr-2">🎵</span>MelodyCritic
            </h3>
            <p className="text-gray-400">
              Платформа для настоящих ценителей музыки. Оценивайте, обсуждайте и открывайте новые музыкальные жемчужины.
            </p>
           
          </div>

          {/* Навигация */}
          <div>
            <h4 className="text-lg font-semibold text-white mb-4">Навигация</h4>
            <ul className="space-y-3">
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">Главная</a></li>
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">Топ-100</a></li>
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">Рецензии</a></li>
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">Жанры</a></li>
            </ul>
          </div>

          {/* Сообщество */}
          <div>
            <h4 className="text-lg font-semibold text-white mb-4">Сообщество</h4>
            <ul className="space-y-3">
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">Форумы</a></li>
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">Рейтинги</a></li>
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">События</a></li>
              <li><a href="#" className="text-gray-400 hover:text-purple-400 transition-colors">Блог</a></li>
            </ul>
          </div>

          {/* Контакты и подписка */}
          <div className="space-y-4">
            <h4 className="text-lg font-semibold text-white">Подписаться на новости</h4>
            <p className="text-gray-400">
              Получайте уведомления о новых релизах и событиях первыми.
            </p>
            <form className="flex">
              <input
                type="email"
                placeholder="Ваш email"
                className="px-4 py-2 rounded-l-md bg-gray-800 text-white focus:outline-none focus:ring-2 focus:ring-purple-500 w-full placeholder-gray-500"
              />
              <button
                type="submit"
                className="bg-gradient-to-r from-purple-500 to-pink-400 hover:from-purple-600 hover:to-pink-500 px-4 py-2 rounded-r-md text-white font-medium transition-all duration-200"
              >
                OK
              </button>
            </form>
              
            </div>
          </div>
        </div>

        {/* Нижняя часть футера */}
        <div className="border-t border-gray-800 mt-12 pt-8 flex flex-col md:flex-row justify-between items-center">
          <p className="text-gray-500 text-sm">
            © {new Date().getFullYear()} MelodyCritic. Все права защищены.
          </p>
          <div className="flex space-x-6 mt-4 md:mt-0">
            <a href="#" className="text-gray-500 hover:text-purple-400 text-sm transition-colors">Политика конфиденциальности</a>
            <a href="#" className="text-gray-500 hover:text-purple-400 text-sm transition-colors">Условия использования</a>
            <a href="#" className="text-gray-500 hover:text-purple-400 text-sm transition-colors">Контакты</a>
          </div>
        </div>
    </footer>
  )
}

export { Footer }