
-- Hailstone Length

hailstone n
    | even n = div n 2
    | odd n = 3 * n + 1
hailLen 1 = 0
hailLen x
    | x>2^64 = -1
    | otherwise = hailLen (hailstone x) + 1
    
    
-- Primes and Divisors
    

divisors :: Int -> [Int]
divisors n = [ i | i <- [2..(n `div` 2)], n `mod` i == 0]
primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == [] ]


-- Joining Strings

join :: [Char] -> [[Char]] -> [Char]
join x y
    | length y ==0 = []
    | length y == 1 = y !! 0
    | length y > 1 = join x (((y !! 0) ++ x ++ (y !! 1)) : drop 2 y)  


-- Pythagorean Triples

pythagorean :: Int -> [(Int, Int, Int)]
pythagorean n = [(a,b,c)|a<-[1..n],b<-[1..n],c<-[5..n],a^2+b^2==c^2,a<=b]