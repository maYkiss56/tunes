import { useState, useEffect, useMemo } from "react";
import { useGenres } from "./useGenres";
import { useTracksRating } from "./useTracksRating";
import type { Track } from "../types";

export const useTrackFilters = (initialTracks: Track[]) => {
  const { genres, loading: genresLoading } = useGenres();
  const { tracks: ratedTracks, loading: ratingLoading } = useTracksRating();
  const [searchQuery, setSearchQuery] = useState("");
  const [debouncedQuery, setDebouncedQuery] = useState("");
  const [selectedGenre, setSelectedGenre] = useState("all");
  const [sortOption, setSortOption] = useState("popular");
  const [isFiltering, setIsFiltering] = useState(false);

  // Добавляем debounce для поиска
  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedQuery(searchQuery);
      setIsFiltering(false);
    }, 300);

    return () => {
      clearTimeout(timer);
      setIsFiltering(true);
    };
  }, [searchQuery]);

  const ratingMap = useMemo(() => {
    const map = new Map<number, number>();
    ratedTracks.forEach((track) => {
      map.set(track.id, track.rating || 0);
    });
    return map;
  }, [ratedTracks]);

  // Функция сортировки
  const sortTracks = (tracks: Track[], option: string) => {
    switch (option) {
      case "newest":
        return [...tracks].sort(
          (a, b) =>
            new Date(b.release_date).getTime() -
            new Date(a.release_date).getTime(),
        );
      case "popular":
        return [...tracks].sort((a, b) => {
          const ratingA = ratingMap.get(a.id) || 0;
          const ratingB = ratingMap.get(b.id) || 0;
          return ratingB - ratingA;
        });
      case "name":
        return [...tracks].sort((a, b) => a.title.localeCompare(b.title));
      default:
        return tracks;
    }
  };

  const filteredTracks = useMemo(() => {
    let result = [...initialTracks];

    // Поиск
    if (debouncedQuery) {
      const query = debouncedQuery.toLowerCase();
      result = result.filter(
        (track) =>
          track.title.toLowerCase().includes(query) ||
          track.artist.nickname.toLowerCase().includes(query) ||
          track.genre?.title.toLowerCase().includes(query),
      );
    }

    // Фильтрация по жанру
    if (selectedGenre !== "all") {
      result = result.filter(
        (track) => track.genre?.id.toString() === selectedGenre,
      );
    }

    // Сортировка
    return sortTracks(result, sortOption);
  }, [initialTracks, debouncedQuery, selectedGenre, sortOption, ratingMap]);

  // Мемоизированные опции
  const genreOptions = useMemo(
    () => [
      { value: "all", label: "Все жанры" },
      ...(genres?.map((genre) => ({
        value: genre.id.toString(),
        label: genre.title,
      })) || []),
    ],
    [genres],
  );

  const sortOptions = useMemo(
    () => [
      { value: "popular", label: "По популярности" },
      { value: "newest", label: "По новизне" },
      { value: "name", label: "По названию" },
    ],
    [],
  );

  return {
    filteredTracks,
    isFiltering: isFiltering || ratingLoading,
    searchQuery,
    setSearchQuery,
    selectedGenre,
    setSelectedGenre,
    sortOption,
    setSortOption,
    genreOptions,
    sortOptions,
    genresLoading: genresLoading || ratingLoading,
  };
};
