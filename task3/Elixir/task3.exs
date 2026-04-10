defmodule Main do
  # проверка на полный квадрат
  def is_square(x) do
    root = :math.sqrt(x) |> trunc()
    root * root == x
  end

  def main do
    IO.puts("1 - input from console\n2 - input from file")
    {choice, _} = Integer.parse(IO.gets("") |> String.trim())

    count =
      cond do
        choice == 1 ->
          IO.puts("Enter N:")
          {n, _} = Integer.parse(IO.gets("") |> String.trim())

          Enum.reduce(1..n, 0, fn _, acc ->
            {x, _} = Integer.parse(IO.gets("") |> String.trim())
            if is_square(x), do: acc + 1, else: acc
          end)

        choice == 2 ->
          IO.puts("Enter file name:")
          filename = IO.gets("") |> String.trim()

          case File.read(filename) do
            {:ok, content} ->
              nums =
                content
                |> String.split()
                |> Enum.map(&String.to_integer/1)

              [n | rest] = nums

              rest
              |> Enum.take(n)
              |> Enum.count(&is_square/1)

            _ ->
              IO.puts("File open error")
              0
          end

        true ->
          IO.puts("Invalid choice")
          0
      end

    IO.puts("Result: #{count}")
  end
end

Main.main()
