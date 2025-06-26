import { Link } from 'react-router-dom'

const NotFoundPage = () => {
	return (
		<div className='min-h-screen bg-gray-900 text-white'>
			<main className='flex flex-col items-center justify-center py-20 px-4 text-center'>
				<div className='max-w-2xl mx-auto'>
					{/* Анимированный номер 404 */}
					<div className='text-9xl font-bold mb-6 bg-gradient-to-r from-purple-500 to-pink-600 bg-clip-text text-transparent'>
						404
					</div>

					<h1 className='text-4xl md:text-5xl font-bold mb-6'>
						Упс! Страница потерялась в музыке
					</h1>

					<p className='text-xl text-gray-300 mb-8'>
						Кажется, вы пытаетесь найти что-то, чего нет в нашей коллекции.
						Возможно, страница была удалена или вы ошиблись адресом.
					</p>

					<div className='flex flex-col sm:flex-row justify-center gap-4'>
						<Link
							to='/'
							className='bg-purple-600 hover:bg-purple-700 text-white font-bold py-3 px-6 rounded-lg transition duration-300'
						>
							На главную
						</Link>

						<Link
							to='/songs'
							className='bg-gray-800 hover:bg-gray-700 text-white font-bold py-3 px-6 rounded-lg transition duration-300'
						>
							К списку песен
						</Link>
					</div>

					{/* Декоративный элемент */}
					<div className='mt-12 relative'>
						<div className='absolute inset-0 flex items-center justify-center'>
							<div className='w-50 h-50 rounded-full bg-purple-900 opacity-20 animate-pulse'></div>
						</div>
						<svg
							className='w-30 h-30 mx-auto relative z-10'
							fill='none'
							stroke='currentColor'
							viewBox='0 0 24 24'
							xmlns='http://www.w3.org/2000/svg'
						>
							<path
								strokeLinecap='round'
								strokeLinejoin='round'
								strokeWidth='1.5'
								d='M9 19V6l12-3v13M9 19c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zm12-3c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zM9 10l12-3'
							></path>
						</svg>
					</div>
				</div>
			</main>
		</div>
	)
}

export default NotFoundPage
