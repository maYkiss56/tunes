import { type FC, type ReactNode } from 'react'

type ButtonProps = {
  children: ReactNode
  width?: number | string
  height?: number | string
  className?: string
  variant?: 'primary' | 'secondary' | 'danger' | 'success'
  size?: 'sm' | 'md' | 'lg'
  onClick?: () => void
  disabled?: boolean
  type?: 'button' | 'submit' | 'reset'
}

const Button: FC<ButtonProps> = ({
  children,
  width,
  height,
  className = '',
  variant = 'primary',
  size = 'md',
  onClick,
  disabled = false,
  type = 'button',
  ...props
}) => {
  // Базовые стили
  const baseStyles = 'rounded-md font-medium transition-all duration-200 focus:outline-none focus:ring-2 cursor-pointer'
  
  // Варианты кнопки
  const variants = {
    primary: 'bg-purple-600 hover:bg-purple-700 text-white focus:ring-purple-300',
    secondary: 'bg-gray-200 hover:bg-gray-300 text-gray-800 focus:ring-gray-300',
    danger: 'bg-red-500 hover:bg-red-600 text-white focus:ring-red-300',
    success: 'bg-green-500 hover:bg-green-600 text-white focus:ring-green-300'
  }
  
  // Размеры кнопки
  const sizes = {
    sm: 'py-1 px-3 text-sm',
    md: 'py-2 px-4 text-base',
    lg: 'py-3 px-6 text-lg'
  }
  
  // Инлайн стили для ширины и высоты
  const inlineStyles = {
    ...(width && { width: typeof width === 'number' ? `${width}px` : width }),
    ...(height && { height: typeof height === 'number' ? `${height}px` : height })
  }
  
  return (
    <button
      className={`${baseStyles} ${variants[variant]} ${sizes[size]} ${className}`}
      style={inlineStyles}
      onClick={onClick}
      disabled={disabled}
      type={type}
      {...props}
    >
      {children}
    </button>
  )
}

export { Button }




