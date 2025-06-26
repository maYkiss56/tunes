import { useEffect, useState } from "react";
import { useAuth } from "../context/AuthContext";
import type { Review } from "../types";

export const useProfileReviews = () => {
  const [reviews, setReviews] = useState<Review[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const { user, isAuthenticated } = useAuth();

  useEffect(() => {
    const fetchReviews = async () => {
      if (!isAuthenticated || !user) {
        setLoading(false);
        return;
      }

      try {
        setLoading(true);
        setError(null);

        const res = await fetch(
          `http://localhost:8080/api/reviews/user/${user.id}`,
          {
            credentials: "include",
          },
        );

        const data: Review[] = await res.json();
        setReviews(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Неизвестная ошибка");
        console.error("Ошибка загрузки рецензий:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchReviews();
  }, [user, isAuthenticated]);

  return {
    reviews,
    loading,
    error,
    isAuthenticated,
  };
};
