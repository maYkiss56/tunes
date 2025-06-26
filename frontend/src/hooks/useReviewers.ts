import { useState, useEffect } from "react";
import type { Reviewer } from "../types";

export const useReviewers = () => {
  const [reviewers, setReviewers] = useState<Reviewer[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchReviewers = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/reviewers");
        if (!response.ok) {
          throw new Error("Failed to fetch reviewers");
        }
        const data = await response.json();
        setReviewers(data);
      } catch (err) {
        console.log(err);
      } finally {
        setLoading(false);
      }
    };

    fetchReviewers();
  }, []);

  return { reviewers, loading };
};
