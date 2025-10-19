import React, { useState, useEffect } from 'react';
import { useLocation, useParams, useNavigate } from 'react-router-dom';

const EditBookPage = () => {
  const { id } = useParams();
  const location = useLocation();
  const navigate = useNavigate();
  const [book, setBook] = useState(location.state?.book || null);
  const [loading, setLoading] = useState(!book);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!book) {
      // fetch book data
      const fetchBook = async () => {
        try {
          setLoading(true);
          const res = await fetch(`http://localhost:8080/api/v1/books/${id}`);
          if (!res.ok) throw new Error('Failed to fetch book');
          const data = await res.json();
          setBook(data);
        } catch (err) {
          setError(err.message);
        } finally {
          setLoading(false);
        }
      };

      fetchBook();
    }
  }, [id, book]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`http://localhost:8080/api/v1/books/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(book),
      });
      if (!res.ok) throw new Error('Failed to update book');
      navigate('/allbooks');
    } catch (err) {
      alert('Update failed: ' + err.message);
    }
  };

  if (loading) return <div className="p-8">Loading...</div>;
  if (error) return <div className="p-8 text-red-600">Error: {error}</div>;

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-2xl font-bold mb-4">แก้ไขหนังสือ</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium">Title</label>
          <input value={book.title || ''} onChange={(e) => setBook({...book, title: e.target.value})} className="mt-1 block w-full border rounded p-2" />
        </div>
        <div>
          <label className="block text-sm font-medium">Author</label>
          <input value={book.author || ''} onChange={(e) => setBook({...book, author: e.target.value})} className="mt-1 block w-full border rounded p-2" />
        </div>
        <div className="flex gap-2">
          <button type="submit" className="px-4 py-2 bg-viridian-600 text-white rounded">Save</button>
          <button type="button" onClick={() => navigate(-1)} className="px-4 py-2 border rounded">Cancel</button>
        </div>
      </form>
    </div>
  );
};

export default EditBookPage;
