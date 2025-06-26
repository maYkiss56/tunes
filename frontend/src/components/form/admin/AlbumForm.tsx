import { useEffect, useRef, useState, type ChangeEvent } from "react";
import type { Album, Artist } from "../../../types";
import { Button } from "../../ui/Button";
import _exports from "tailwind-scrollbar";

interface AlbumFormProps {
  initialData: Album | null;
  onSubmit: (data: Album, imageFile?: File) => void;
  onCancel: () => void;
}

const AlbumForm = ({ onCancel, initialData, onSubmit }: AlbumFormProps) => {
  const [title, setTitle] = useState(initialData?.title || "");
  const [imagePreview, setImagePreview] = useState("");
  const [imageFile, setImageFile] = useState<File | null>(null);
  const [artistID, setArtistID] = useState(initialData?.artist.id || 0);
  const [artists, setArtists] = useState<Artist[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/api/artists")
      .then((res) => res.json())
      .then(setArtists);
  }, []);

  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    setTitle(initialData?.title || "");
    setArtistID(initialData?.artist.id || 0);
    setImageFile(null);

    const preview = initialData?.image_url
      ? initialData.image_url.startsWith("http")
        ? initialData.image_url
        : `http://localhost:8080/${initialData.image_url}`
      : "";

    setImagePreview(preview);
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

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const selectedArtist = artists.find((a) => a.id === artistID);

    if (!selectedArtist) {
      return;
    }

    onSubmit(
      {
        id: initialData?.id || 0,
        title,
        image_url: imagePreview,
        artist: selectedArtist,
      },
      imageFile ?? undefined,
    );
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Title {!initialData && <span className="text-red-500">*</span>}
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
        :q
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Image {!initialData && <span className="text-red-500">*</span>}
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

export default AlbumForm;
