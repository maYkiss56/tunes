import { useEffect, useRef, useState, type ChangeEvent } from "react";
import type { Genre } from "../../../types";
import { Button } from "../../ui/Button";

interface GenreFormProps {
  initialData: Genre | null;
  onSubmit: (data: Genre, imageFile?: File) => void;
  onCancel: () => void;
}

const GenreForm = ({ initialData, onSubmit, onCancel }: GenreFormProps) => {
  const [title, setTitle] = useState(initialData?.title || "");
  const [imagePreview, setImagePreview] = useState("");
  const [imageFile, setImageFile] = useState<File | null>(null);

  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    setTitle(initialData?.title || "");
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
    onSubmit(
      {
        id: initialData?.id || 0,
        title,
        image_url: imagePreview,
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

export default GenreForm;
