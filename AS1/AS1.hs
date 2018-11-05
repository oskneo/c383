import RainbowAssign
import qualified Data.Map as Map
import qualified Data.Maybe as Maybe
import Data.Int

pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table

-- Create the function to convert hash integer to base 2 integer and divide the integer into list using div and mod 
hash2baseNletter :: Hash -> [Int]
hash2baseNletter hs = toN (fromEnum hs) []
    where
        toN :: Int -> [Int] -> [Int]
        toN hs' xs
            | hs' <= 0 && hs' > -nLetters = xs
            | otherwise = toN (div hs' nLetters) ([mod hs' nLetters] ++ xs)


-- Create a function to take significant numbers
takesign :: [Int] -> [Int]
takesign xs
    | length xs > pwLength = takesign (tail xs)
    | length xs < pwLength = takesign ([0] ++ xs)
    | otherwise = xs

            
-- Create a function convert integers to letter starting from a using succ
int2letter :: [Char] -> Int -> [Char]
int2letter xs n = xs ++ [nxLetter n 'a']
    where
        nxLetter :: Int -> Char -> Char
        nxLetter n' x'
            | n' == 0 = x'
            | otherwise = nxLetter (n' - 1) (succ x')
            

-- Create a reduced function by converting hash value to base2 interger then change them to letters
pwReduce :: Hash -> Passwd
pwReduce n = foldl int2letter "" (takesign (hash2baseNletter n))

-- A simple function for do a hash and reduce function
circle :: Hash -> Hash
circle h = pwHash $ pwReduce h


-- Create a rainbowtable and map the table
rainbowTable :: Int -> [Passwd] -> Map.Map Hash Passwd
rainbowTable n xs = Map.fromList $ rt xs
    where
        rainbow :: Passwd -> [Hash]
        rainbow h2 = iterate circle (pwHash h2)
        tp :: Passwd -> (Hash,Passwd)
        tp h3= ((rainbow h3) !! n, h3)
        rt :: [Passwd] -> [(Hash,Passwd)]
        rt xs' = map tp xs' 
        

generateTable :: IO ()
generateTable = do
  table <- buildTable rainbowTable nLetters pwLength width height
  writeTable table filename
  
  
-- Create a function for finding password finally by recursively applying hash and reduced function to hash to find the correct row then find the corresponding password.
findPassword :: Map.Map Hash Passwd -> Int -> Int32 -> Maybe Passwd
findPassword table n hash = findPW n  hash
    where 
        searchrow :: Passwd -> Int -> Hash -> Maybe Passwd
        searchrow pw' n' h'
            | n' == -1 = Nothing
            | h' == hash = Just pw'
            | otherwise = searchrow (pwReduce h') (n'-1) $ circle h'
        checkn :: Hash -> Int -> Maybe Passwd
        checkn h' n'
            | n' == 0 = Nothing
            | otherwise = findPW (n'-1)  $ circle h'
        findPW :: Int -> Hash -> Maybe Passwd
        findPW n' h = case Map.lookup h table of
                    Nothing -> checkn h n'
                    Just y  -> searchrow y n $ pwHash y

test1 :: IO (Maybe Passwd)
test1 = do
  table <- readTable filename
  return (Map.lookup 0 table)
  

test2 :: Int -> IO ([Passwd], Int)
test2 n = do
  table <- readTable filename
  pws <- randomPasswords nLetters pwLength n
  let hs = map pwHash pws
  let result = Maybe.mapMaybe (findPassword table width) hs
  return (result, length result)
  
  
main :: IO ()
main = do
  generateTable
  res <- test2 1000
  print res