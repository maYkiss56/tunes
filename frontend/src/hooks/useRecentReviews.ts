// hooks/useRecentReviews.ts
import { useEffect, useState } from "react";
import type { Review } from "../types";

export const useRecentReviews = () => {
  const [recentReviews, setRecentReviews] = useState<Review[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchReviews = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/reviews");
        const data: Review[] = await res.json();
        // отсортировать по дате
        const sorted = data.sort(
          (a, b) =>
            new Date(b.updated_at).getTime() - new Date(a.updated_at).getTime(),
        );
        setRecentReviews(sorted.slice(0, 6));
      } catch (e) {
        console.error("Ошибка загрузки рецензий", e);
      } finally {
        setLoading(false);
      }
    };

    fetchReviews();
  }, []);

  return { recentReviews, loading };
};
