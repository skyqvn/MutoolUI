unit MainUnit;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, Menus, ComCtrls,
  StdCtrls, ExtCtrls, Buttons, Grids, AnchorDockPanel;

type

  { TMainForm }

  TMainForm = class(TForm)
    MainMenu: TMainMenu;
    Panel1: TPanel;
    Panel2: TPanel;
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

end.
