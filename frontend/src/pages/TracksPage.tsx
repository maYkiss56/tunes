import { useState, useEffect } from "react";
import type { Track } from "../types";
import { useTrackFilters } from "../hooks/useTrackFilters";
import { SearchInput } from "../components/ui/SearchInput";
import { Select } from "../components/ui/Select";
import { TrackCard } from "../components/cards/TrackCard";
import ModelWindow from "../components/blocks/ModelWindow";
import { Spinner } from "../components/ui/Spinner";
import { FilterIcon } from "../components/ui/icons";
import { Header } from "../components/blocks/Header";
import { Footer } from "../components/blocks/Footer";

const TracksPage = () => {
  const [initialTracks, setInitialTracks] = useState<Track[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedTrack, setSelectedTrack] = useState<Track | null>(null);
  const [showMobileFilters, setShowMobileFilters] = useState(false);

  const {
    filteredTracks,
    isFiltering,
    searchQuery,
    setSearchQuery,
    selectedGenre,
    setSelectedGenre,
    sortOption,
    setSortOption,
    genreOptions,
    sortOptions,
    genresLoading,
  } = useTrackFilters(initialTracks);

  useEffect(() => {
    const fetchTracks = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/songs");
        if (!res.ok) throw new Error("Ошибка при загрузке песен");
        const data = await res.json();
        setInitialTracks(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Неизвестная ошибка");
      } finally {
        setLoading(false);
      }
    };

    fetchTracks();
  }, []);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <Spinner size="lg" />
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center text-red-500 py-8">
        Ошибка загрузки: {error}
      </div>
    );
  }

  return (
    <>
      <Header />
      <div className="bg-gray-900 min-h-screen text-white">
        {/* Фильтры и поиск */}
        <div className="sticky top-0 z-10 bg-gray-900/95 backdrop-blur-sm py-4 px-4 border-b border-gray-800 w-304 mx-auto">
          <div className="max-w-7xl mx-auto">
            <div className="flex flex-col gap-4">
              {/* Мобильный фильтр */}
              <button
                onClick={() => setShowMobileFilters(!showMobileFilters)}
                className="md:hidden flex items-center gap-2 bg-gray-800 px-4 py-2 rounded-lg"
              >
                <FilterIcon className="w-5 h-5" />
                Фильтры
              </button>

              <div
                className={`${showMobileFilters ? "block" : "hidden"} md:block`}
              >
                <div className="flex flex-col md:flex-row gap-4 items-center justify-between">
                  <div className="w-full md:w-1/2">
                    <SearchInput
                      value={searchQuery}
                      onChange={setSearchQuery}
                      placeholder="Поиск треков, исполнителей..."
                    />
                  </div>

                  <div className="flex gap-3 w-full md:w-auto">
                    <Select
                      value={selectedGenre}
                      onChange={(e) => setSelectedGenre(e.target.value)}
                      options={genreOptions}
                      className="flex-1"
                      disabled={genresLoading}
                    />
                    <Select
                      value={sortOption}
                      onChange={(e) => setSortOption(e.target.value)}
                      options={sortOptions}
                      className="flex-1 "
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Индикатор фильтрации */}
        {isFiltering && (
          <div className="fixed top-4 right-4 bg-gray-800 px-3 py-1 rounded-full text-sm flex items-center gap-2 z-20">
            <Spinner size="sm" />
            Фильтрация...
          </div>
        )}

        {/* Список треков */}
        <div className="max-w-7xl mx-auto px-4 py-8">
          {filteredTracks.length > 0 ? (
            <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-6">
              {filteredTracks.map((track) => (
                <TrackCard
                  key={track.id}
                  track={track}
                  onClick={() => setSelectedTrack(track)}
                />
              ))}
            </div>
          ) : (
            <div className="text-center py-12 text-gray-400">
              Ничего не найдено. Попробуйте изменить параметры поиска.
            </div>
          )}
        </div>

        {/* Модальное окно с деталями трека */}
        {selectedTrack && (
          <ModelWindow
            track={selectedTrack}
            onClose={() => setSelectedTrack(null)}
          />
        )}
      </div>
      <Footer />
    </>
  );
};

export default TracksPage;
