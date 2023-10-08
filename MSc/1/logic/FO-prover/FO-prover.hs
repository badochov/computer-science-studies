{-# LANGUAGE FlexibleInstances #-}
{-# LANGUAGE LambdaCase #-}
{-# LANGUAGE TypeSynonymInstances #-}
{-# LANGUAGE UnicodeSyntax #-}

module Main where

import Control.Monad.State
import Data.Containers.ListUtils (nubOrd)
import Data.List
import qualified Data.Map
import qualified Data.Set
import Debug.Trace
import Formula
import Parser hiding (one)
import System.IO
import System.Random
import System.Timeout
import Test.QuickCheck hiding (Fun, (===))
import Text.Printf
import Utils

trace' s a = if False then trace s a else a

prover :: Formula -> Bool
prover _ = False

-- lab 2

type RProp = (RelName, [Term])

data Literal = Pos RProp | Neg RProp deriving (Eq, Show, Ord)

literal2var :: Literal -> RProp
literal2var (Pos p) = p
literal2var (Neg p) = p

opposite :: Literal -> Literal
opposite (Pos p) = Neg p
opposite (Neg p) = Pos p

type CNFClause = [Literal]

type CNF = [CNFClause]

cnf2formula :: CNF -> Formula
cnf2formula [] = T
cnf2formula lss = foldr1 And (map go lss)
  where
    go [] = F
    go ls = foldr1 Or (map go2 ls)
    go2 (Pos (n, ts)) = Rel n ts
    go2 (Neg (n, ts)) = Not (Rel n ts)

positiveLiterals :: CNFClause -> [RProp]
positiveLiterals ls = nubOrd [p | Pos p <- ls]

negativeLiterals :: CNFClause -> [RProp]
negativeLiterals ls = nubOrd [p | Neg p <- ls]

literals :: [Literal] -> [RProp]
literals ls = nubOrd $ positiveLiterals ls ++ negativeLiterals ls

cnf :: Formula -> CNF
cnf = go . nnf
  where
    go T = []
    go F = [[]]
    go (Rel n ts) = [[Pos (n, ts)]]
    go (Not (Rel n ts)) = [[Neg (n, ts)]]
    go (φ `And` ψ) = go φ ++ go ψ
    go (φ `Or` ψ) = [as ++ bs | as <- go φ, bs <- go ψ]
    go _ = error "Impossible"

ecnf :: Formula -> CNF
ecnf f =
  let (c, ff, _) = ecnf' f $ rprops f
   in cnf ff ++ c
  where
    rprops :: Formula -> [RProp]
    rprops T = []
    rprops F = []
    rprops (Rel n ts) = [(n, ts)]
    rprops (Not phi) = rprops phi
    rprops (And phi psi) = nubOrd $ rprops phi ++ rprops psi
    rprops (Or phi psi) = nubOrd $ rprops phi ++ rprops psi
    rprops (Implies phi psi) = nubOrd $ rprops phi ++ rprops psi
    rprops (Iff phi psi) = nubOrd $ rprops phi ++ rprops psi
    rprops (Exists _ phi) = rprops phi
    rprops (Forall _ phi) = rprops phi
    ecnf'' :: Formula -> Formula -> [RProp] -> (Formula -> Formula -> Formula) -> (CNF, Formula, [RProp])
    ecnf'' phi psi vs fc =
      let (cnf_phi, phi', vs') = ecnf' phi vs
          (cnf_psi, psi', vs'') = ecnf' psi vs'
          (vn, vts) = freshRProp vs''
          cnf_iff = cnf $ Iff (Rel vn vts) (fc phi' psi')
       in (cnf_iff ++ (cnf_phi ++ cnf_psi), Rel vn vts, (vn, vts) : vs'')
    freshRProp :: [RProp] -> RProp
    freshRProp vars = head $ filter (`notElem` vars) $ map (\x -> ("p" ++ show x, [])) [0 ..]
    ecnf' :: Formula -> [RProp] -> (CNF, Formula, [RProp])
    ecnf' T vs = ([], T, vs)
    ecnf' F vs = ([], F, vs)
    ecnf' (Rel p ts) vs = ([], Rel p ts, vs)
    ecnf' (Not phi) vs =
      let (cnf_phi, phi', vs') = ecnf' phi vs
       in case phi' of
            T -> (cnf_phi, F, vs')
            F -> (cnf_phi, T, vs')
            Rel p ts -> (cnf_phi, Not $ Rel p ts, vs')
            Not (Rel p ts) -> (cnf_phi, Rel p ts, vs')
            _ -> error "impossible"
    ecnf' (And phi psi) vs = ecnf'' phi psi vs And
    ecnf' (Or phi psi) vs = ecnf'' phi psi vs Or
    ecnf' (Implies phi psi) vs = ecnf'' phi psi vs Implies
    ecnf' (Iff phi psi) vs = ecnf'' phi psi vs Iff
    ecnf' _ _ = error "impossible"

removeTautologies :: CNF -> CNF
removeTautologies c = go c []
  where
    is_tautology :: CNFClause -> Bool
    is_tautology cs = is_tautology' cs Data.Set.empty Data.Set.empty
      where
        is_tautology' [] _ _ = False
        is_tautology' ((Pos p) : t) pos neg = Data.Set.member p neg || is_tautology' t (Data.Set.insert p pos) neg
        is_tautology' ((Neg p) : t) pos neg = Data.Set.member p pos || is_tautology' t pos (Data.Set.insert p neg)
    go :: CNF -> CNF -> CNF
    go [] res = res
    go (h : t) res =
      let res' = if is_tautology h then res else h : res
       in go t res'

oneLiteral :: CNF -> CNF
oneLiteral cnf =
  let res = foldl' reduce cnf $ gather cnf
   in if res == cnf then res else oneLiteral res
  where
    reduce :: CNF -> Literal -> CNF
    reduce cnf l = reduce' cnf []
      where
        reduce' [] res = res
        reduce' (h : t) res =
          reduce' t $
            if l `elem` h
              then res
              else filter (opposite l /=) h : res
    gather cnf = gather' cnf []
      where
        gather' [] res = res
        gather' ([a] : t) res = gather' t (a : res)
        gather' (h : t) res = gather' t res

affirmativeNegative :: CNF -> CNF
affirmativeNegative cnf =
  let res = filter (not . has_any_literal find_removable) cnf
   in if res == cnf then res else affirmativeNegative res
  where
    has_any_literal :: [RProp] -> CNFClause -> Bool
    has_any_literal s clause = any (`elem` literals clause) s
    find_removable :: [RProp]
    find_removable =
      let (pos, neg) = foldl' (\(p, n) c -> (positiveLiterals c ++ p, negativeLiterals c ++ n)) ([], []) cnf
          pos_set = Data.Set.fromList pos
          neg_set = Data.Set.fromList neg
          intersection = Data.Set.intersection pos_set neg_set
       in Data.Set.elems $ Data.Set.difference (Data.Set.union pos_set neg_set) intersection

resolution :: CNF -> CNF
resolution f =
  case least_used_var of
    Nothing -> f
    Just v ->
      let (with_v, without_v) = split_formulas v f
          v_t = sub (Pos v) with_v []
          v_f = sub (Neg v) with_v []
          f_t = cnf2formula v_t
          f_f = cnf2formula v_f
          cnf_res = cnf $ Or f_f f_t
       in nubOrd $ cnf_res ++ without_v
  where
    least_used_var :: Maybe RProp
    least_used_var =
      let counts = get_counts f
       in snd <$> Data.Map.foldlWithKey min' Nothing counts
      where
        min' Nothing v c = Just (c, v)
        min' (Just prev) v c = Just $ min prev (c, v)
        get_counts = foldl' get_counts_clause Data.Map.empty
        get_counts_clause = foldl' (\m l -> Data.Map.alter addOne (literal2var l) m)
        addOne Nothing = Just 1
        addOne (Just a) = Just (a + 1)
    sub _ [] res = res
    sub p (h : t) res =
      sub p t $
        if p `elem` h
          then res
          else filter (opposite p /=) h : res
    split_formulas v =
      foldl'
        ( \(with, without) h ->
            if v `elem` literals h
              then (h : with, without)
              else (with, h : without)
        )
        ([], [])

dp :: CNF -> Bool
dp [] = True
dp cnf =
  [] `notElem` cnf
    && (dp . resolution . affirmativeNegative . oneLiteral . removeTautologies) cnf

satDP :: SATSolver
satDP form = dp cnf
  where
    cnf = ecnf form

-- lab 3

varsT :: Term -> [VarName]
varsT (Var x) = [x]
varsT (Fun _ ts) = nubOrd $ concatMap varsT ts

vars :: Formula -> [VarName]
vars T = []
vars F = []
vars (Rel _ ts) = varsT (Fun "dummy" ts)
vars (Not phi) = vars phi
vars (And phi psi) = nubOrd $ vars phi ++ vars psi
vars (Or phi psi) = nubOrd $ vars phi ++ vars psi
vars (Implies phi psi) = nubOrd $ vars phi ++ vars psi
vars (Iff phi psi) = nubOrd $ vars phi ++ vars psi
vars (Exists x phi) = nubOrd $ x : vars phi
vars (Forall x phi) = nubOrd $ x : vars phi

freshIn :: VarName -> Formula -> Bool
x `freshIn` phi = x `notElem` vars phi

variants :: VarName -> [VarName]
variants x = x : [x ++ show n | n <- [0 ..]]

freshVariant :: VarName -> Formula -> VarName
freshVariant x phi = head [y | y <- variants x, y `freshIn` phi]

nnf :: Formula -> Formula
nnf T = T
nnf F = F
nnf rel@(Rel _ _) = rel
nnf (And phi psi) = And (nnf phi) (nnf psi)
nnf (Or phi psi) = Or (nnf phi) (nnf psi)
nnf (Implies phi psi) = Or (nnf $ Not phi) (nnf psi)
nnf (Iff phi psi) = And (Or (nnf $ Not phi) (nnf psi)) (Or (nnf $ Not psi) (nnf phi))
nnf (Exists v phi) = Exists v (nnf phi)
nnf (Forall v phi) = Forall v (nnf phi)
nnf (Not T) = F
nnf (Not F) = T
nnf (Not rel@(Rel _ _)) = Not $ nnf rel
nnf (Not (And phi psi)) = Or (nnf $ Not phi) (nnf $ Not psi)
nnf (Not (Or phi psi)) = And (nnf $ Not phi) (nnf $ Not psi)
nnf (Not (Implies phi psi)) = And (nnf phi) (nnf $ Not psi)
nnf (Not (Iff phi psi)) = Or (And (nnf phi) (nnf $ Not psi)) (And (nnf psi) (nnf $ Not phi))
nnf (Not (Exists v phi)) = Forall v (nnf $ Not phi)
nnf (Not (Forall v phi)) = Exists v (nnf $ Not phi)
nnf (Not (Not a)) = nnf a

renameT :: VarName -> VarName -> Term -> Term
renameT x y (Var z)
  | z == x = Var y
  | otherwise = Var z
renameT x y (Fun f ts) = Fun f (map (renameT x y) ts)

rename :: VarName -> VarName -> Formula -> Formula
rename x y T = T
rename x y F = F
rename x y (Rel r ts) = Rel r (map (renameT x y) ts)
rename x y (Not phi) = Not (rename x y phi)
rename x y (And phi psi) = And (rename x y phi) (rename x y psi)
rename x y (Or phi psi) = Or (rename x y phi) (rename x y psi)
rename x y (Implies phi psi) = Implies (rename x y phi) (rename x y psi)
rename x y (Iff phi psi) = Iff (rename x y phi) (rename x y psi)
rename x y (Forall z phi)
  | z == x = Forall z phi
  | otherwise = Forall z (rename x y phi)
rename x y (Exists z phi)
  | z == x = Exists z phi
  | otherwise = Exists z (rename x y phi)

-- lab 4

skolemise :: Formula -> Formula
skolemise = pnf . replaceExistential . miniscope . uniqueExistential . nnf . closeFormula
  where
    closeFormula :: Formula -> Formula
    -- Formula is already closed by universal quantifiers.
    closeFormula = id

    uniqueExistential :: Formula -> Formula
    uniqueExistential phi = evalState (go phi) $ fv phi
      where
        go :: Formula -> State [VarName] Formula
        go T = return T
        go F = return F
        go phi@(Rel _ _) = return phi
        go (Not phi) = Not <$> go phi
        go (And phi psi) = liftM2 And (go phi) (go psi)
        go (Or phi psi) = liftM2 Or (go phi) (go psi)
        go (Implies phi psi) = liftM2 Implies (go phi) (go psi)
        go (Iff phi psi) = liftM2 Iff (go phi) (go psi)
        go (Forall x phi) = go2 Forall x phi
        go (Exists x phi) = go2 Exists x phi

        go2 quantifier x phi =
          do
            xs <- get
            let y = head [y | y <- variants x, y `notElem` xs]
            let psi = rename x y phi
            put $ y : xs
            quantifier y <$> go psi

    miniscope :: Formula -> Formula
    miniscope T = T
    miniscope F = F
    miniscope r@(Rel _ _) = r
    miniscope (Not r@(Rel _ _)) = Not r
    miniscope (And phi psi) = And (miniscope phi) (miniscope psi)
    miniscope (Or phi psi) = Or (miniscope phi) (miniscope psi)
    miniscope (Exists x phi) = case miniscope phi of
            T -> Exists x T
            F -> Exists x F
            r@(Rel _ _) -> Exists x r
            Not r@(Rel _ _) -> Exists x $ Not r
            And phi psi -> miniscope_bi Exists x And phi psi
            Or phi psi -> miniscope_bi Exists x Or phi psi
            Exists x' phi -> Exists x' $ miniscope $ Exists x phi
            Forall x' phi -> Forall x' $ miniscope $ Exists x phi
            _ -> error "impossible"
    miniscope (Forall x phi) = case miniscope phi of
            T -> Forall x T
            F -> Forall x F
            r@(Rel _ _) -> Forall x r
            Not r@(Rel _ _) -> Forall x $ Not r
            And phi psi -> miniscope_bi Forall x And phi psi
            Or phi psi -> miniscope_bi Forall x Or phi psi
            Exists x' phi -> Forall x $ Exists x' phi
            Forall x' phi -> Forall x' $ miniscope $ Forall x phi
            _ -> error "impossible"
    miniscope _ = error "Impossible"
    miniscope_bi q x op phi psi =
          let inPhi = x `elem` fv phi
              inPsi = x `elem` fv psi
           in if inPhi && inPsi
                then q x (op (miniscope phi) (miniscope psi))
                else
                  if inPhi
                    then op (miniscope $ q x phi) (miniscope psi)
                    else op (miniscope phi) (miniscope $ q x psi)

    replaceExistential :: Formula -> Formula
    replaceExistential phi = r phi Data.Map.empty []
      where
        -- TODO ensure there's no name clash
        to_skolem_function :: Data.Map.Map VarName [Term] -> Term -> Term
        to_skolem_function m (Var x) = case Data.Map.lookup x m of
          Just vs -> Fun x vs
          Nothing -> Var x
        to_skolem_function m (Fun n ts) = Fun n $ map (to_skolem_function m) ts
        r :: Formula -> Data.Map.Map VarName [Term] -> [Term] -> Formula
        r T _ _ = T
        r F _ _ = F
        r (Rel n ts) m _ = Rel n $ map (to_skolem_function m) ts
        r (Not phi) m vs = Not $ r phi m vs
        r (And phi psi) m vs = And (r phi m vs) (r psi m vs)
        r (Or phi psi) m vs = Or (r phi m vs) (r psi m vs)
        r (Exists x phi) m vs = r phi (Data.Map.insert x vs m) vs
        r (Forall x phi) m vs = Forall x $ r phi m (Var x : vs)
        r _ _ _ = error "Impossible"

    pnf :: Formula -> Formula
    pnf phi = phi

atomicFormulas :: Formula -> [Formula]
atomicFormulas T = []
atomicFormulas F = []
atomicFormulas phi@(Rel _ ts) = [phi]
atomicFormulas (Not phi) = atomicFormulas phi
atomicFormulas (And phi psi) = nubOrd $ atomicFormulas phi ++ atomicFormulas psi
atomicFormulas (Or phi psi) = nubOrd $ atomicFormulas phi ++ atomicFormulas psi
atomicFormulas (Implies phi psi) = nubOrd $ atomicFormulas phi ++ atomicFormulas psi
atomicFormulas (Iff phi psi) = nubOrd $ atomicFormulas phi ++ atomicFormulas psi
atomicFormulas (Exists x phi) = atomicFormulas phi
atomicFormulas (Forall x phi) = atomicFormulas phi

groundInstances :: Formula -> [Term] -> [Formula]
groundInstances phi ts =
  let subs = replicateM (length free_vars) ts
   in map (\s -> apply (make_sub s) phi) subs
  where
    free_vars = fv phi
    make_sub s = (Data.Map.!) (Data.Map.fromList $ zip free_vars s)

-- sat :: Formula -> Bool
-- sat phi = or [ev int phi | int <- functions atoms [True, False]]
--   where
--     atoms = atomicFormulas phi

--     ev :: (Formula -> Bool) -> Formula -> Bool
--     ev int T = True
--     ev int F = False
--     ev int atom@(Rel _ _) = int atom
--     ev int (Not φ) = not (ev int φ)
--     ev int (Or φ ψ) = ev int φ || ev int ψ
--     ev int (And φ ψ) = ev int φ && ev int ψ
--     ev _ φ = error $ "unexpected formula: " ++ show φ

type Arity = Int

type Signature = [(FunName, Arity)]

sigT :: Term -> Signature
sigT (Var _) = []
sigT (Fun f ts) = nubOrd $ (f, length ts) : concatMap sigT ts

sig :: Formula -> Signature
sig T = []
sig F = []
sig (Rel r ts) = concatMap sigT ts
sig (Not phi) = sig phi
sig (And phi psi) = nubOrd $ sig phi ++ sig psi
sig (Or phi psi) = nubOrd $ sig phi ++ sig psi
sig (Implies phi psi) = nubOrd $ sig phi ++ sig psi
sig (Iff phi psi) = nubOrd $ sig phi ++ sig psi
sig (Exists _ phi) = sig phi
sig (Forall _ phi) = sig phi

constants :: Signature -> [Term]
constants s = [Fun c [] | (c, 0) <- s]

funs :: Signature -> [(FunName, Arity)]
funs s = [(c, x) | (c, x) <- s, x > 0]

generalise :: Formula -> Formula
generalise phi = go $ fv phi
  where
    go [] = phi
    go (x : xs) = Forall x $ go xs

aedecide :: Formula -> Bool
aedecide phi =
  let f = (remove_forall . skolemise . Not . generalise) phi
      universum = generate_herbrands_universum f
      gis = groundInstances f universum
      prefixes = map (foldl1' And) $ tail $ inits gis
   in not $ all satDP prefixes
  where
    remove_forall (Forall _ phi) = remove_forall phi
    remove_forall T = T
    remove_forall F = F
    remove_forall r@(Rel _ _) = r
    remove_forall (Not phi) = Not $ remove_forall phi
    remove_forall (And phi psi) = And (remove_forall phi) (remove_forall psi)
    remove_forall (Or phi psi) = Or (remove_forall phi) (remove_forall psi)
    remove_forall _ = error "impossible"
    consts s = case constants s of
      [] -> [Fun "" []]
      ts -> ts
    generate' :: [Term] -> [(FunName, Arity)] -> [Term]
    generate' _ [] = []
    generate' cs fns = [Fun name args | depth <- [0 ..], (name, arity) <- fns, args <- replicateM arity $ generate'' depth cs fns]
    generate'' :: Int -> [Term] -> [(FunName, Arity)] -> [Term]
    generate'' 0 cs _ = cs
    generate'' x cs fns = generate'' (x -1) cs fns ++ [Fun name args | (name, arity) <- fns, args <- replicateM arity $ generate'' (x -1) cs fns]
    generate_herbrands_universum f =
      let s = sig f
          cs = consts s
       in cs ++ generate' cs (funs s)

-- main

main :: IO ()
main = do
  eof <- isEOF
  if eof
    then return ()
    else do
      line <- getLine -- read the input
      let phi = parseString line -- call the parser
      let res = aedecide phi -- call the prover
      if res
        then putStrLn "1" -- write 1 if the formula is a tautology
        else putStrLn "0" -- write 0 if the formula is not a tautology