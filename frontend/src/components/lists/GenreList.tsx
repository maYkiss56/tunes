import { useEffect, useState } from "react";
import type { Genre } from "../../types";
import { GenreCard } from "../cards/GenreCard";

const GenreList = () => {
  const [genres, setGenres] = useState<Genre[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchGenres = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/genres");
        if (!res.ok) throw new Error("Ошибка при загрузке жанров");
        const data = await res.json();
        setGenres(data);
      } catch (error) {
        console.error("Ошибка при получении жанров:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchGenres();
  }, []);

  if (loading)
    return <p className="text-center text-white">Загрузка жанров...</p>;

  return (
    <section className="w-304 mx-auto py-16 px-4">
      <h2 className="text-3xl font-bold text-center mb-12">Жанры</h2>
      <div className="flex overflow-x-auto gap-6 pb-4 scrollbar scrollbar-thin scrollbar-thumb-purple-500 scrollbar-track-transparent">
        {genres.map((genre) => (
          <div key={genre.id} className="flex-shrink-0 cursor-pointer">
            <GenreCard genre={genre} />
          </div>
        ))}
      </div>
    </section>
  );
};

export { GenreList };
