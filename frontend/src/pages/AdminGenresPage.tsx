import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import DataTable, { type Column } from "../components/blocks/admin/DataTable";
import GenreForm from "../components/form/admin/GenreForm";
import { Button } from "../components/ui/Button";
import { PlusIcon } from "../components/ui/icons";
import type { Genre } from "../types";

const AdminGenresPage = () => {
  const [genres, setGenres] = useState<Genre[]>([]);
  const [currentGenre, setCurrentGenre] = useState<Genre | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();
  const { id } = useParams();

  const isFormOpen = !!id;

  const onCancel = () => {
    navigate("/admin/genres");
  };

  useEffect(() => {
    const fetchGenres = async () => {
      setIsLoading(true);
      try {
        const res = await fetch("http://localhost:8080/api/genres", {
          method: "GET",
          credentials: "include",
        });
        const data = await res.json();
        if (Array.isArray(data)) {
          setGenres(data);
        }
      } catch (error) {
        console.error("Error fetching genres:", error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchGenres();
  }, []);

  useEffect(() => {
    if (id && id !== "new") {
      const numericId = Number(id);
      const fetchGenre = async () => {
        try {
          const res = await fetch(
            `http://localhost:8080/api/genres/${numericId}`,
          );
          const data = await res.json();
          setCurrentGenre(data);
        } catch (error) {
          console.error("Error fetching genre:", error);
        }
      };
      fetchGenre();
    } else if (id === "new") {
      setCurrentGenre(null);
    }
  }, [id]);

  const handleEdit = (genre: Genre) => {
    navigate(`/admin/genres/${genre.id}`);
  };

  const handleDelete = async (id: number) => {
    try {
      const res = await fetch(`http://localhost:8080/api/admin/genres/${id}`, {
        method: "DELETE",
        credentials: "include",
      });

      if (!res.ok) throw new Error("Failed to delete genre");

      setGenres((prev) => prev.filter((genre) => genre.id !== id));
    } catch (err) {
      console.error("Error deleting genre:", err);
    }
  };

  const handleSubmit = async (genreData: Genre, imageFile?: File) => {
    const formData = new FormData();
    formData.append("title", genreData.title);

    if (imageFile) {
      formData.append("image", imageFile);
    }

    try {
      let res: Response;
      let url: string;
      let method: string;

      if (currentGenre) {
        url = `http://localhost:8080/api/admin/genres/${currentGenre.id}`;
        method = "PATCH";
      } else {
        url = "http://localhost:8080/api/admin/genres";
        method = "POST";
      }

      // eslint-disable-next-line prefer-const
      res = await fetch(url, {
        method,
        credentials: "include",
        body: formData,
      });

      if (!res.ok) {
        const errorData = await res.json();
        throw new Error(errorData.message || "Request failed");
      }

      const updatedGenre = await res.json();

      if (currentGenre) {
        setGenres((prev) =>
          prev.map((genre) =>
            genre.id === currentGenre.id ? updatedGenre : genre,
          ),
        );
      } else {
        const res = await fetch("http://localhost:8080/api/genres", {
          method: "GET",
          credentials: "include",
        });
        const data = await res.json();
        if (Array.isArray(data)) {
          setGenres(data);
        }
      }

      navigate("/admin/genres");
    } catch (err) {
      console.error("Error submitting genre:", err);
    }
  };

  const columns: Column<Genre>[] = [
    { key: "id", header: "ID" },
    { key: "title", header: "Title" },
    {
      key: "image_url",
      header: "Image",
      render: (value, _item) => (
        <img
          src={
            typeof value === "string"
              ? value.startsWith("http")
                ? value
                : `http://localhost:8080/${value}`
              : undefined
          }
          alt="Genre"
          className="h-10 w-10 object-cover rounded"
        />
      ),
    },
  ];

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">Manage Genres</h2>
        <Button onClick={() => navigate("/admin/genres/new")} variant="primary">
          <PlusIcon className="w-5 h-5 mr-1" />
          Добавить жанр
        </Button>
      </div>

      {isFormOpen ? (
        <div className="bg-white p-6 rounded-lg shadow-sm mb-6">
          <GenreForm
            initialData={currentGenre}
            onSubmit={handleSubmit}
            onCancel={onCancel}
          />
        </div>
      ) : (
        <DataTable<Genre>
          columns={columns}
          data={genres}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      )}
    </div>
  );
};

export default AdminGenresPage;
