import { ru } from "date-fns/locale/ru";
import { useEffect, useRef, useState, type ChangeEvent } from "react";
import DatePicker, { registerLocale } from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import type { Album, Artist, Genre, Track } from "../../../types";
import { Button } from "../../ui/Button";

registerLocale("ru", ru);

interface SongFormProps {
  initialData: Track | null;
  onSubmit: (data: Track, imageFile?: File) => void;
  onCancel: () => void;
}

const SongForm = ({ onCancel, initialData, onSubmit }: SongFormProps) => {
  const [title, setTitle] = useState(initialData?.title || "");
  const [imagePreview, setImagePreview] = useState(
    initialData?.image_url
      ? initialData.image_url.startsWith("http")
        ? initialData.image_url
        : `http://localhost:8080/${initialData.image_url}`
      : "",
  );
  const [fullTitle, setFullTitle] = useState(initialData?.full_title || "");
  const [imageFile, setImageFile] = useState<File | null>(null);
  const [releaseDate, setReleaseDate] = useState<Date>(
    initialData?.release_date ? new Date(initialData.release_date) : new Date(),
  );

  const [genreID, setGenreID] = useState(initialData?.genre.id || 0);
  const [artistID, setArtistID] = useState(initialData?.artist.id || 0);
  const [albumID, setAlbumID] = useState(initialData?.album?.id || 0);
  const [genres, setGenres] = useState<Genre[]>([]);
  const [artists, setArtists] = useState<Artist[]>([]);
  const [albums, setAlbums] = useState<Album[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/api/artists")
      .then((res) => res.json())
      .then(setArtists);

    fetch("http://localhost:8080/api/albums")
      .then((res) => res.json())
      .then(setAlbums);

    fetch("http://localhost:8080/api/genres")
      .then((res) => res.json())
      .then(setGenres);
  }, []);

  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    setTitle(initialData?.title || "");
    setFullTitle(initialData?.full_title || "");
    setImagePreview(
      initialData?.image_url
        ? initialData.image_url.startsWith("http")
          ? initialData.image_url
          : `http://localhost:8080/${initialData.image_url}`
        : "",
    );
    setImageFile(null);
    setReleaseDate(
      initialData?.release_date
        ? new Date(initialData.release_date)
        : new Date(),
    );
    setGenreID(initialData?.genre.id || 0);
    setArtistID(initialData?.artist.id || 0);
    setAlbumID(initialData?.album?.id || 0);
  }, [initialData]);

  const handleImageChange = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setImageFile(file);
      const reader = new FileReader();
      reader.onload = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleTitleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const handleFullTitleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setFullTitle(e.target.value);
  };

  const handleReleaseDateChange = (date: Date | null) => {
    if (date) {
      setReleaseDate(date);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (
      !initialData &&
      (!title ||
        !fullTitle ||
        !releaseDate ||
        !imageFile ||
        !genreID ||
        !artistID ||
        !albumID)
    ) {
      alert("Please fill all required fields");
      return;
    }

    const songData: Track = {
      id: initialData?.id || 0,
      title,
      full_title: fullTitle,
      image_url: imagePreview.startsWith("http://localhost:8080/")
        ? imagePreview.replace("http://localhost:8080/", "")
        : imagePreview,
      release_date: releaseDate.toISOString(),
      genre: genres.find((g) => g.id === genreID)!,
      artist: artists.find((a) => a.id === artistID)!,
      album: albumID ? albums.find((a) => a.id === albumID) : undefined,
    };
    onSubmit(songData, imageFile || undefined);
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Название {!initialData && <span className="text-red-500">*</span>}
          </label>
          <input
            type="text"
            name="title"
            value={title}
            onChange={handleTitleChange}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 text-gray-700"
            required={!initialData}
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Полное название{" "}
            {!initialData && <span className="text-red-500">*</span>}
          </label>
          <input
            type="text"
            name="full_title"
            value={fullTitle}
            onChange={handleFullTitleChange}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 text-gray-700"
            required={!initialData}
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Обложка {!initialData && <span className="text-red-500">*</span>}
          </label>
          <input
            type="file"
            name="image_url"
            accept="image/*"
            onChange={handleImageChange}
            ref={fileInputRef}
            className="hidden"
          />
          <Button
            type="button"
            variant="secondary"
            onClick={() => fileInputRef.current?.click()}
            className="mb-2"
          >
            {imageFile ? "Change Image" : "Upload Image"}
          </Button>
          {imagePreview && (
            <div className="mt-2">
              <img
                src={
                  imagePreview.startsWith("data:") ||
                  imagePreview.startsWith("http")
                    ? imagePreview
                    : `http://localhost:8080${imagePreview}`
                }
                alt="Preview"
                className="h-32 w-32 object-cover rounded"
              />
            </div>
          )}
        </div>

        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700">
            Дата релиза{" "}
            {!initialData && <span className="text-red-500">*</span>}
          </label>
          <DatePicker
            selected={releaseDate}
            onChange={handleReleaseDateChange}
            dateFormat="yyyy-MM-dd"
            locale="ru"
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-600 text-gray-700"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Жанр {!initialData && <span className="text-red-500">*</span>}
          </label>
          <div className="relative">
            <select
              className="w-full px-4 py-2 border border-gray-300 rounded-lg bg-white text-gray-700 focus:outline-none focus:ring-2 focus:ring-purple-600 appearance-none"
              name="artist_id"
              value={genreID}
              onChange={(e) => setGenreID(Number(e.target.value))}
              required={!initialData}
            >
              <option value={0}>Выберите жанр</option>
              {genres.map((genre) => (
                <option key={genre.id} value={genre.id}>
                  {genre.title}
                </option>
              ))}
            </select>
            <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
              <svg
                className="fill-current h-4 w-4"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
              >
                <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
              </svg>
            </div>
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Артист {!initialData && <span className="text-red-500">*</span>}
          </label>
          <div className="relative">
            <select
              className="w-full px-4 py-2 border border-gray-300 rounded-lg bg-white text-gray-700 focus:outline-none focus:ring-2 focus:ring-purple-600 appearance-none"
              name="artist_id"
              value={artistID}
              onChange={(e) => setArtistID(Number(e.target.value))}
              required={!initialData}
            >
              <option value={0}>Выберите артиста</option>
              {artists.map((artist) => (
                <option key={artist.id} value={artist.id}>
                  {artist.nickname}
                </option>
              ))}
            </select>
            <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
              <svg
                className="fill-current h-4 w-4"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
              >
                <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
              </svg>
            </div>
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Альбом {!initialData && <span className="text-red-500">*</span>}
          </label>
          <div className="relative">
            <select
              className="w-full px-4 py-2 border border-gray-300 rounded-lg bg-white text-gray-700 focus:outline-none focus:ring-2 focus:ring-purple-600 appearance-none"
              name="album_id"
              value={albumID}
              onChange={(e) => setAlbumID(Number(e.target.value))}
              required={!initialData}
            >
              <option value={0}>Выберите альбом</option>
              {albums.map((album) => (
                <option key={album.id} value={album.id}>
                  {album.title}
                </option>
              ))}
            </select>
            <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
              <svg
                className="fill-current h-4 w-4"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
              >
                <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
              </svg>
            </div>
          </div>
        </div>
      </div>

      <div className="flex space-x-2">
        <Button type="submit" variant="primary">
          {initialData ? "Обновить" : "Добавить"}
        </Button>
        <Button type="button" onClick={onCancel} variant="secondary">
          Отмена
        </Button>
      </div>
    </form>
  );
};

export default SongForm;
