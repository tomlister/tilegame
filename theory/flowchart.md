~~~flow
```flow
st=>start: Start
op=>operation: Use Weapon
issword=>condition: Is Item Sword?
iswand=>condition: Is Item Wand?
getitemtype=>operation: Get Item Held
getmouseposition=>inputoutput: Get Mouse Position
getmouseposition2=>inputoutput: Get Mouse Position
checkcollisionwithenemy=>operation: Check Mouse Collision With Enemy
iscoll=>condition: Collided?
damageenemy=>operation: Damage Enemy Health
spawnspell=>parallel: Spawn Spell
movespell=>subroutine: Move Spell
playmisssound=>subroutine: Play Miss Sound
playhitsound=>subroutine: Play Hit Sound
e=>end

st->op->getitemtype->issword
issword(yes)->getmouseposition->checkcollisionwithenemy->iscoll
issword(no)->iswand
iswand(yes)->getmouseposition2->spawnspell
iswand(no)->e
iscoll(yes)->damageenemy->playhitsound->e
iscoll(no)->playmisssound->e
spawnspell(path1, bottom)->movespell
spawnspell(path2, right)->e
```
~~~

