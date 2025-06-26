import { type FC, type ReactNode } from 'react'

type AuthFormProps = {
	title: string
	subtitle: string
	children: ReactNode
}

const AuthForm: FC<AuthFormProps> = ({ title, subtitle, children }) => {
	return (
		<div className='min-h-screen bg-gradient-to-b from-gray-900 to-black flex items-center justify-center p-4'>
			<div className='w-full max-w-md bg-gray-800 rounded-xl shadow-2xl overflow-hidden'>
				<div className='bg-gradient-to-r from-purple-900 to-pink-800 p-6 text-center'>
					<h1 className='text-3xl font-bold text-white'>{title}</h1>
					<p className='text-purple-200 mt-2'>{subtitle}</p>
				</div>
				<div className='p-6 space-y-6'>{children}</div>
			</div>
		</div>
	)
}

export { AuthForm }
