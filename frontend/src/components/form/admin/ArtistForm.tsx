import { useEffect, useState, type ChangeEvent } from 'react'
import type { Artist } from '../../../types'
import { Button } from '../../ui/Button'

interface ArtistFormProps {
	initialData: Artist | null
	onSubmit: (data: Artist) => void
	onCancel: () => void
}

const ArtistForm = ({ onCancel, initialData, onSubmit }: ArtistFormProps) => {
	const [nickname, setNickname] = useState(initialData?.nickname || '')
	const [bio, setBio] = useState(initialData?.bio || '')
	const [country, setCountry] = useState(initialData?.country || '')

	useEffect(() => {
		setNickname(initialData?.nickname || '')
		setBio(initialData?.bio || '')
		setCountry(initialData?.country || '')
	}, [initialData])

	const handleNicknameChange = (e: ChangeEvent<HTMLInputElement>) =>
		setNickname(e.target.value)
	const handleBioChange = (e: ChangeEvent<HTMLTextAreaElement>) =>
		setBio(e.target.value)
	const handleCountryChange = (e: ChangeEvent<HTMLInputElement>) =>
		setCountry(e.target.value)

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault()

		if (!initialData && (!nickname || !bio || !country)) {
			alert('Пожалуйста, заполните все обязательные поля')
			return
		}

		const artistData: Artist = {
			id: initialData?.id || 0,
			nickname,
			bio,
			country,
		}
		onSubmit(artistData)
	}

	return (
		<form onSubmit={handleSubmit} className='space-y-6'>
			<div className='grid grid-cols-1 md:grid-cols-2 gap-6'>
				<div className='md:col-span-2'>
					<label className='block text-sm font-medium text-gray-700 mb-1'>
						Nickname {!initialData && <span className='text-red-500'>*</span>}
					</label>
					<input
						type='text'
						value={nickname}
						onChange={handleNicknameChange}
						className='w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 text-gray-700'
						required={!initialData}
					/>
				</div>

				<div className='md:col-span-2'>
					<label className='block text-sm font-medium text-gray-700 mb-1'>
						Biography {!initialData && <span className='text-red-500'>*</span>}
					</label>
					<textarea
						value={bio}
						onChange={handleBioChange}
						className='w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 text-gray-700'
						rows={4}
						required={!initialData}
					/>
				</div>

				<div className='md:col-span-2'>
					<label className='block text-sm font-medium text-gray-700 mb-1'>
						Country {!initialData && <span className='text-red-500'>*</span>}
					</label>
					<input
						type='text'
						value={country}
						onChange={handleCountryChange}
						className='w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 text-gray-700'
						required={!initialData}
					/>
				</div>
			</div>

			<div className='flex justify-end space-x-3 pt-4'>
				<Button type='button' variant='secondary' onClick={onCancel}>
					Назад
				</Button>
				<Button type='submit'>
					{initialData?.id ? 'Обновить' : 'Добавить'}
				</Button>
			</div>
		</form>
	)
}

export default ArtistForm
