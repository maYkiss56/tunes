import { type FC } from "react";
import { Button } from "../ui/Button";
import { Link } from "react-router-dom";
import { useAuth } from "../../context/AuthContext";

const Header: FC = () => {
  const { user, logout, loading } = useAuth();

  if (loading) {
    return null; //TODO: skeleton
  }

  const isAuthenticated = !!user;

  return (
    <header className="bg-gradient-to-r from-gray-900 to-black text-white sticky top-0 z-50 shadow-lg">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-20">
          <div className="flex items-center space-x-10">
            <Link to="/" className="flex-shrink-0 flex items-center">
              <span className="text-2xl mr-2">🎵</span>
              <h1 className="text-2xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-300">
                MelodyCritic
              </h1>
            </Link>
            <nav className="hidden md:flex space-x-8">
              <Link
                to="/"
                className="text-gray-300 hover:text-purple-400 px-3 py-2 text-md font-medium transition-colors duration-200"
              >
                Домой
              </Link>
              <Link
                to="/tracks"
                className="text-gray-300 hover:text-purple-400 px-3 py-2 text-md font-medium transition-colors duration-200"
              >
                Обзор музыки
              </Link>
              <Link
                to="/top"
                className="text-gray-300 hover:text-purple-400 px-3 py-2 text-md font-medium transition-colors duration-200"
              >
                Топ рейтингов
              </Link>
            </nav>
          </div>

          <div className="flex items-center space-x-4">
            {isAuthenticated ? (
              <>
                <Button
                  variant="secondary"
                  size="md"
                  onClick={logout}
                  className="bg-gray-700 hover:bg-gray-600 text-white"
                >
                  Выйти
                </Button>
                <Link to="/profile">
                  <Button
                    size="md"
                    className="hidden md:inline-flex bg-gradient-to-r from-purple-500 to-pink-400 hover:from-purple-600 hover:to-pink-500"
                  >
                    Профиль
                  </Button>
                </Link>
              </>
            ) : (
              <>
                <Link to="/login">
                  <Button
                    variant="secondary"
                    size="md"
                    className="hidden md:inline-flex bg-gray-700 hover:bg-gray-600 text-white"
                  >
                    Войти
                  </Button>
                </Link>
                <Link to="/register">
                  <Button
                    size="md"
                    className="bg-gradient-to-r from-purple-500 to-pink-400 hover:from-purple-600 hover:to-pink-500"
                  >
                    Регистрация
                  </Button>
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </header>
  );
};

export { Header };
