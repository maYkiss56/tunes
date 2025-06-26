import { useEffect, useState } from "react";
import type { Genre } from "../types";

export const useGenres = () => {
  const [genres, setGenres] = useState<Genre[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchGenres = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/genres");
        const data: Genre[] = await res.json();
        setGenres(data);
      } catch (e) {
        console.log("Ошибка загрузки жанров", e);
      } finally {
        setLoading(false);
      }
    };
    fetchGenres();
  }, []);

  return { genres, loading };
};
